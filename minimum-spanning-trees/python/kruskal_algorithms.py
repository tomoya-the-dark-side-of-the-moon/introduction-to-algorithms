from disjoint_set import Node, find_set, union_set
from graph import Edge, Graph, Vertex


def mst_kruskal(graph):
    vertices = graph.vertices
    nodes = {v.name: Node(v.name) for v in vertices.values()}

    edges = graph.edges
    edges.sort()

    results = []
    for edge in edges:
        u, v, weight = edge.start, edge.finish, edge.weight
        if find_set(nodes[u]) != find_set(nodes[v]):
            print(f"add edge {u, v} ({weight})")

            union_set(nodes[u], nodes[v])
            results.append(edge)

    return results


if __name__ == "__main__":
    vertices = {name: Vertex(name) for name in "abcdefghi"}

    edges = []
    edges.append(Edge("a", "b", weight=4))
    edges.append(Edge("a", "h", weight=8))
    edges.append(Edge("b", "a", weight=4))
    edges.append(Edge("b", "c", weight=8))
    edges.append(Edge("b", "h", weight=11))
    edges.append(Edge("c", "b", weight=8))
    edges.append(Edge("c", "d", weight=7))
    edges.append(Edge("c", "f", weight=4))
    edges.append(Edge("c", "i", weight=2))
    edges.append(Edge("d", "c", weight=7))
    edges.append(Edge("d", "e", weight=9))
    edges.append(Edge("d", "f", weight=14))
    edges.append(Edge("e", "d", weight=9))
    edges.append(Edge("e", "f", weight=10))
    edges.append(Edge("f", "c", weight=4))
    edges.append(Edge("f", "d", weight=14))
    edges.append(Edge("f", "e", weight=10))
    edges.append(Edge("f", "g", weight=2))
    edges.append(Edge("g", "f", weight=2))
    edges.append(Edge("g", "h", weight=1))
    edges.append(Edge("g", "i", weight=6))
    edges.append(Edge("h", "a", weight=8))
    edges.append(Edge("h", "b", weight=11))
    edges.append(Edge("h", "g", weight=1))
    edges.append(Edge("h", "i", weight=7))
    edges.append(Edge("i", "c", weight=2))
    edges.append(Edge("i", "g", weight=6))
    edges.append(Edge("i", "h", weight=7))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Kruskal's algorithm ---")
    results = mst_kruskal(graph)

    print("--- result ---")
    total = 0
    for edge in results:
        print(f"start = {edge.start}, finish = {edge.finish} / weight = {edge.weight}")
        total += edge.weight
    print(f"total = {total}")
