from collections import defaultdict


class Node:

    __slots__ = ["data", "parent", "rank"]

    def __init__(self, data):
        self.data = data
        self.parent = self
        self.rank = 0


def union_set(x, y):
    link(find_set(x), find_set(y))


def link(x, y):
    if x.rank > y.rank:
        y.parent = x
    else:
        x.parent = y
        if x.rank == y.rank:
            y.rank += 1


def find_set(x):
    if x.parent is not x:
        x.parent = find_set(x.parent)

    return x.parent


def print_sets(nodes):
    groups = defaultdict(list)
    for node in nodes.values():
        groups[find_set(node).data].append(node.data)

    for root, members in groups.items():
        print(f"{root} : {', '.join(members)}")


if __name__ == "__main__":
    nodes = {ch: Node(ch) for ch in "abcdefghij"}

    print("--- disjoint set after initializing ---")
    print_sets(nodes)
    print("--- merge b and d ---")
    union_set(nodes["b"], nodes["d"])
    print_sets(nodes)
    print("--- merge e and f ---")
    union_set(nodes["e"], nodes["f"])
    print_sets(nodes)
    print("--- merge a and c ---")
    union_set(nodes["a"], nodes["c"])
    print_sets(nodes)
    print("--- merge h and i ---")
    union_set(nodes["h"], nodes["i"])
    print_sets(nodes)
    print("--- merge a and b ---")
    union_set(nodes["a"], nodes["b"])
    print_sets(nodes)
    print("--- merge f and g ---")
    union_set(nodes["f"], nodes["g"])
    print_sets(nodes)
    print("--- merge b and c ---")
    union_set(nodes["b"], nodes["c"])
    print_sets(nodes)
