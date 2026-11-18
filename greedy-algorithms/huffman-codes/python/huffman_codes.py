import heapq


class Element:

    def __init__(self, ch, freq=0):
        self.ch = ch
        self.freq = freq
        self.left = self.right = None

    def __repr__(self):
        return f"{self.ch if self.ch != '' else ''}({self.freq})"

    def __lt__(self, other):
        return self.freq < other.freq


def huffman(characters, frequencies):
    elements = []
    for ch, freq in zip(characters, frequencies):
        elements.append(Element(ch, freq=freq))

    heapq.heapify(elements)

    while len(elements) > 1:
        x = heapq.heappop(elements)
        y = heapq.heappop(elements)
        print(x, y)

        z = Element("", freq=x.freq + y.freq)
        z.left = x
        z.right = y

        heapq.heappush(elements, z)

    return heapq.heappop(elements)


def print_tree(root):

    def print_tree_internal(r, space_size=0):
        if r.left:
            print_tree_internal(r.left, space_size=space_size + 1)

        print("    " * space_size, r)

        if r.right:
            print_tree_internal(r.right, space_size=space_size + 1)

    print_tree_internal(root)


if __name__ == "__main__":
    characters = ["a", "b", "c", "d", "e", "f"]
    frequencies = [45, 13, 12, 16, 9, 5]

    print(" ch |", end="")
    for ch in characters:
        print(f"{ch:>3}|", end="")
    print()
    print("freq|", end="")
    for freq in frequencies:
        print(f"{freq:>3}|", end="")
    print()

    root = huffman(characters, frequencies)
    print_tree(root)
