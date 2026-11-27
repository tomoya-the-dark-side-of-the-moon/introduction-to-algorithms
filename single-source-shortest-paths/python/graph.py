import sys

from collections import defaultdict


class Vertex:

    def __init__(self, name):
        self.name = name
        self.distance = 0
        self.predecessor = None


class Edge:

    def __init__(self, start, finish, weight=0):
        self.start = start
        self.finish = finish
        self.weight = weight

    def __lt__(self, other):
        return self.weight < other.weight


class Graph:

    def __init__(self, vertices, edges):
        self.vertices = vertices
        self.edges = edges
        self.adjacency_lists = defaultdict(list)
        self.__initialize_graph()

    def print_graph(self):
        for v in self.vertices.values():
            print(v.name, ": ", end="")
            adjacency_list = self.adjacency_lists[v.name]
            for i, (adjacency, weight) in enumerate(adjacency_list):
                print(
                    f"{adjacency}({weight})",
                    end=" -> " if i < len(adjacency_list) - 1 else "\n",
                )

    def __initialize_graph(self):
        for edge in self.edges:
            self.adjacency_lists[edge.start].append((edge.finish, edge.weight))


def initialize_single_source(graph, s):
    vertices = graph.vertices

    for v in vertices.values():
        v.distance = sys.maxsize
        v.predecessor = None

    s.distance = 0


def relax(u, v, weight):
    if v.distance > u.distance + weight:
        print(f"{u.name, v.name} before = {v.distance}, after = {u.distance + weight}")

        v.distance = u.distance + weight
        v.predecessor = u
