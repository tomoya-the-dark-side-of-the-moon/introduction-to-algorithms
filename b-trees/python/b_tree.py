class Element:

    def __init__(self):
        self.keys = []
        self.n = 0
        self.children = []
        self.is_leaf = True


class Tree:

    def __init__(self, min_degree=2):
        self.root = Element()
        self.min_degree = min_degree

    def insert(self, key):
        r = self.root
        if r.n == 2 * self.min_degree - 1:
            s = Element()
            s.children.insert(0, r)
            s.is_leaf = False

            self.root = s
            self.__split_child(s, 0)

            self.__insert_key_non_full(s, key)
        else:
            self.__insert_key_non_full(r, key)

    def contains(self, key):
        return self.__contains_internal(self.root, key)

    def __contains_internal(self, x, key):
        i = 0
        while i < x.n and x.keys[i] < key:
            i += 1

        if i < x.n and x.keys[i] == key:
            return True
        elif x.is_leaf:
            return False
        else:
            return self.__contains_internal(x.children[i], key)

    def print_tree(self):

        def print_tree_internal(r, space_size=1):
            if not r.n:
                return

            print("  " * space_size, r.keys)

            if not r.is_leaf:
                for child in r.children:
                    print_tree_internal(child, space_size=space_size + 1)

        print_tree_internal(self.root, space_size=0)

    def __insert_key_non_full(self, r, key):
        if r.is_leaf:
            i = r.n - 1
            while i >= 0 and r.keys[i] > key:
                i -= 1

            i += 1

            r.keys.insert(i, key)
            r.n += 1
        else:
            i = r.n - 1
            while i >= 0 and r.keys[i] > key:
                i -= 1

            i += 1

            if r.children[i].n == 2 * self.min_degree - 1:
                self.__split_child(r, i)

                if r.keys[i] < key:
                    i += 1

            self.__insert_key_non_full(r.children[i], key)

    def __split_child(self, x, i):
        y = x.children[i]
        z = Element()

        key_to_up = y.keys[self.min_degree - 1]

        z.keys = y.keys[self.min_degree :]
        z.n = len(z.keys)
        y.keys = y.keys[: self.min_degree - 1]
        y.n = len(y.keys)

        z.is_leaf = y.is_leaf
        if not y.is_leaf:
            z.children = y.children[self.min_degree :]
            y.children = y.children[: self.min_degree]

        x.keys.insert(i, key_to_up)
        x.n += 1
        x.children.insert(i + 1, z)


if __name__ == "__main__":
    t = Tree()

    t.insert(5)
    t.insert(9)
    t.insert(3)
    t.insert(7)
    t.insert(1)
    t.insert(2)
    t.insert(8)
    t.insert(6)
    t.insert(0)
    t.insert(4)

    while True:
        print("1:insert 2:remove 3:print > ", end="")
        op = int(input())

        if op == 1:
            print("input key > ", end="")
            key = int(input())

            t.insert(key)
        elif op == 2:
            print("input key > ", end="")
            key = int(input())

            has_key = t.contains(key)
            print(has_key)
        elif op == 3:
            t.print_tree()
        else:
            print("invalid operation")
