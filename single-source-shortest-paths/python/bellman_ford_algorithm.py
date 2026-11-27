from graph import Edge, Graph, Vertex, initialize_single_source, relax


def bellman_ford(graph, s):
    initialize_single_source(graph, s)

    vertices = graph.vertices
    edges = graph.edges

    for i in range(1, len(vertices)):
        for edge in edges:
            u, v = vertices[edge.start], vertices[edge.finish]
            weight = edge.weight
            relax(u, v, weight)

    for edge in edges:
        u, v = vertices[edge.start], vertices[edge.finish]
        weight = edge.weight
        if v.distance > u.distance + weight:
            return False

    return True


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
    edges.append(Edge("s", "t", weight=6))
    edges.append(Edge("s", "y", weight=7))
    edges.append(Edge("t", "y", weight=8))
    edges.append(Edge("t", "x", weight=5))
    edges.append(Edge("t", "z", weight=-4))
    edges.append(Edge("y", "x", weight=-3))
    edges.append(Edge("y", "z", weight=9))
    edges.append(Edge("x", "t", weight=-2))
    edges.append(Edge("z", "s", weight=2))
    edges.append(Edge("z", "x", weight=7))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Bellman-Ford algorithm ---")
    has_no_negative_weight_cycle = bellman_ford(graph, vertices['s'])

    print("--- result ---")
    if not has_no_negative_weight_cycle:
        print("Graph has a negative-weight cycle.")
    else:
        for v in vertices.values():
            print(f"name = {v.name} / weight = {v.distance} ({create_path(v)})")
