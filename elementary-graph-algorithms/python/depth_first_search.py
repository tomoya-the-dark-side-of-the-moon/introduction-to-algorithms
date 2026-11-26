from graph import Color, Edge, Graph, Vertex

time = 0


def depth_first_search(graph):
    for u in graph.vertices.values():
        u.color = Color.WHITE
        u.predecessor = None

    for u in graph.vertices.values():
        if u.color == Color.WHITE:
            depth_first_search_internal(graph, u)


def depth_first_search_internal(graph, u):
    global time

    time += 1
    u.discovery_time = time
    u.color = Color.GRAY
    print(f"{u.name} is discovered.")

    adjacency_list = graph.adjacency_lists[u.name]
    for adjacency in adjacency_list:
        v = graph.vertices[adjacency]
        if v.color == Color.WHITE:
            v.predecessor = u
            depth_first_search_internal(graph, v)

    time += 1
    u.finish_time = time
    print(f"{u.name} is finished.")


def create_path(v):
    path = [v.name]
    while v.predecessor:
        v = v.predecessor
        path.append(v.name)

    path.reverse()

    return " -> ".join(path)


if __name__ == "__main__":
    vertices = {name: Vertex(name) for name in "uvxywz"}
    edges = []
    edges.append(Edge("u", "v"))
    edges.append(Edge("u", "x"))
    edges.append(Edge("v", "y"))
    edges.append(Edge("x", "v"))
    edges.append(Edge("y", "x"))
    edges.append(Edge("w", "y"))
    edges.append(Edge("w", "z"))
    edges.append(Edge("z", "z"))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- depth first search ---")
    depth_first_search(graph)

    print("--- result ---")
    for u in vertices.values():
        print(f"{u.name} ({u.discovery_time} / {u.finish_time}) : {create_path(u)}")
