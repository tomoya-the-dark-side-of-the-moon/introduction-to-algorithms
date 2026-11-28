import heapq

from graph import Edge, Graph, Vertex, initialize_single_source, relax


def dijkstra(graph, s):
    initialize_single_source(graph, s)

    vertices_list = list(graph.vertices.values())
    heapq.heapify(vertices_list)

    while len(vertices_list):
        u = heapq.heappop(vertices_list)
        adjacency_list = graph.adjacency_lists[u.name]
        for finish, weight in adjacency_list:
            v = graph.vertices[finish]
            relax(u, v, weight)

        heapq.heapify(vertices_list)

def create_path(v):
    path = []
    while v:
        path.append(v.name)

        v = v.predecessor

    path.reverse()

    return " -> ".join(path)


if __name__ == '__main__':
    vertices = {name: Vertex(name) for name in "styxz"}

    edges = []
    edges.append(Edge("s", "t", weight=10))
    edges.append(Edge("s", "y", weight=5))
    edges.append(Edge("t", "y", weight=2))
    edges.append(Edge("t", "x", weight=1))
    edges.append(Edge("y", "t", weight=3))
    edges.append(Edge("y", "x", weight=9))
    edges.append(Edge("y", "z", weight=2))
    edges.append(Edge("x", "z", weight=4))
    edges.append(Edge("z", "s", weight=7))
    edges.append(Edge("z", "x", weight=6))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Dijkstra's algorithm ---")
    dijkstra(graph, vertices["s"])

    print("--- result ---")
    for v in vertices.values():
        print(f"name = {v.name} / weight = {v.distance} ({create_path(v)})")
