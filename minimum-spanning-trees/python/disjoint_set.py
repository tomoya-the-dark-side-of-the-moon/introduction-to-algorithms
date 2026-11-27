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
