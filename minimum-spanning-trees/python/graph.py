from collections import defaultdict


class Vertex:

    def __init__(self, name):
        self.name = name
        self.weight = 0
        self.predecessor = None

    def __lt__(self, other):
        return self.weight < other.weight


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
