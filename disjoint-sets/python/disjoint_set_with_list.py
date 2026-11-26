class Element:

    __slots__ = ["data", "next"]

    def __init__(self, data):
        self.data = data
        self.next = None


class LinkedList:

    __slots__ = ["head", "tail"]

    def __init__(self):
        self.head = self.tail = None

    def append(self, z):
        if not self.head:
            self.head = z
        else:
            self.tail.next = z
        self.tail = z

    def print_list(self):
        p = self.head
        while p:
            print(p.data, end=" -> " if p.next else "\n")

            p = p.next

    def contains(self, data):
        p = self.head
        while p:
            if p.data == data:
                return True

            p = p.next

        return False


def union_set(sets, x, y):
    head_of_x = find_set(sets, x)
    head_of_y = find_set(sets, y)

    if not head_of_x or not head_of_y:
        return

    if head_of_x == head_of_y:
        return

    ll_of_x = sets[head_of_x.data]
    ll_of_y = sets[head_of_y.data]
    if head_of_x.data < head_of_y.data:
        ll_of_x.tail.next = ll_of_y.head
        ll_of_x.tail = ll_of_y.tail

        del sets[head_of_y.data]
    else:
        ll_of_y.tail.next = ll_of_x.head
        ll_of_y.tail = ll_of_x.tail

        del sets[head_of_x.data]


def find_set(sets, data):
    for ll in sets.values():
        if ll.contains(data):
            return ll.head

    return None


def print_set(sets):
    for k, ll in sets.items():
        print(k, ": ", end="")
        ll.print_list()


if __name__ == "__main__":
    sets = {}

    for ch in "abcdefghij":
        ll = LinkedList()
        e = Element(ch)
        ll.append(e)
        sets[ch] = ll

    print("--- disjoint set after initializing ---")
    print_set(sets)

    print("--- merge b and d ---")
    union_set(sets, "b", "d")
    print_set(sets)
    print("--- merge e and f ---")
    union_set(sets, "e", "f")
    print_set(sets)
    print("--- merge a and c ---")
    union_set(sets, "a", "c")
    print_set(sets)
    print("--- merge h and i ---")
    union_set(sets, "h", "i")
    print_set(sets)
    print("--- merge a and b ---")
    union_set(sets, "a", "b")
    print_set(sets)
    print("--- merge f and g ---")
    union_set(sets, "f", "g")
    print_set(sets)
    print("--- merge b and c ---")
    union_set(sets, "b", "c")
    print_set(sets)
