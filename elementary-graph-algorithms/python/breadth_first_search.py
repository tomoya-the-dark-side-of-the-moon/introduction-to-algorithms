from collections import deque
from graph import Color, Edge, Graph, Vertex
import math


def breadth_first_search(graph, s):
    for v in graph.vertices.values():
        v.color, v.distance, v.predecessor = Color.WHITE, math.inf, None
    s.color, s.distance = Color.GRAY, 0
    print(s.name, "color white -> gray")

    queue = deque([s])

    while len(queue) != 0:
        u = queue.popleft()
        adjacency_list = graph.adjacency_lists[u.name]
        for adjacency in adjacency_list:
            v = vertices[adjacency]
            if v.color == Color.WHITE:
                v.color = Color.GRAY
                print(v.name, "color white -> gray")
                print(v.name, f"distance {v.distance} -> {u.distance + 1}")
                v.distance = u.distance + 1
                v.predecessor = u
                print(v.name, "predecessor ->", u.name)

                queue.append(v)

        u.color = Color.BLACK
        print(u.name, "color gray -> black")


def create_path(v):
    path = [v.name]
    while v.predecessor:
        v = v.predecessor
        path.append(v.name)

    path.reverse()

    return " -> ".join(path)


if __name__ == "__main__":
    vertices = {name: Vertex(name) for name in "srvtuywxz"}
    edges = []
    edges.append(Edge("s", "r"))
    edges.append(Edge("s", "v"))
    edges.append(Edge("s", "u"))
    edges.append(Edge("r", "s"))
    edges.append(Edge("r", "t"))
    edges.append(Edge("r", "w"))
    edges.append(Edge("v", "s"))
    edges.append(Edge("v", "y"))
    edges.append(Edge("v", "w"))
    edges.append(Edge("t", "r"))
    edges.append(Edge("t", "u"))
    edges.append(Edge("u", "s"))
    edges.append(Edge("u", "t"))
    edges.append(Edge("u", "y"))
    edges.append(Edge("y", "v"))
    edges.append(Edge("y", "u"))
    edges.append(Edge("y", "x"))
    edges.append(Edge("w", "r"))
    edges.append(Edge("w", "v"))
    edges.append(Edge("w", "x"))
    edges.append(Edge("w", "z"))
    edges.append(Edge("x", "y"))
    edges.append(Edge("x", "w"))
    edges.append(Edge("x", "z"))
    edges.append(Edge("z", "w"))
    edges.append(Edge("z", "x"))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- breadth first search ---")
    breadth_first_search(graph, vertices["s"])

    print("--- result ---")
    for v in vertices.values():
        print(f"{v.name} ({v.distance}) :", create_path(v))
