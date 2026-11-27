import heapq
import sys

from graph import Edge, Graph, Vertex


def mst_prim(graph, r):
    vertices = graph.vertices
    for v in vertices.values():
        v.weight = sys.maxsize
        v.predecessor = None

    r.weight = 0
    vertices_list = list(graph.vertices.values())
    heapq.heapify(vertices_list)

    while len(vertices_list):
        u = heapq.heappop(vertices_list)
        print(f"target vertex: {u.name}, weight = {u.weight}")

        adjacency_list = graph.adjacency_lists[u.name]
        for v_name, weight in adjacency_list:
            v = vertices[v_name]
            if v in vertices_list and v.weight > weight:
                print(f"Update vertex: {v.name}, weight: {v.weight} -> {weight}")
                v.predecessor = u
                v.weight = weight

                heapq.heapify(vertices_list)


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

    print("--- Prim's algorithm ---")
    mst_prim(graph, vertices["a"])

    print("--- result ---")
    total = 0
    for v in vertices.values():
        print(f"name = {v.name}, key = {v.weight}", end="")
        if v.predecessor:
            print(f", predecessor = {v.predecessor.name}", end="")
        print()

        total += v.weight
    print(f"total = {total}")
