from collections import defaultdict
import enum


class Color(enum.Enum):

    WHITE = "white"
    GRAY = "gray"
    BLACK = "black"


class Vertex:

    def __init__(self, name, color=Color.WHITE):
        self.name = name
        self.color = color
        self.distance = 0
        self.predecessor = None


class Edge:

    def __init__(self, start, finish):
        self.start = start
        self.finish = finish


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
            for i, adjacency in enumerate(adjacency_list):
                print(adjacency, end=" -> " if i < len(adjacency_list) - 1 else "\n")

    def __initialize_graph(self):
        for edge in self.edges:
            self.adjacency_lists[edge.start].append(edge.finish)
