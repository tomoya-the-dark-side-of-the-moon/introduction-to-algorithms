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

    def remove(self, key):
        self.__remove_internal(self.root, key)

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

    def __remove_internal(self, x, key):
        i = 0
        while i < x.n and x.keys[i] < key:
            i += 1

        if i == x.n:
            if x.children[i].n >= self.min_degree:
                self.__remove_internal(x.children[i], key)
            else:
                if x.children[i - 1].n >= self.min_degree:
                    self.__sink_key_to_left_node(x, key, i)
                else:
                    self.__merge_key_to_left_node(x, key, i)
        else:
            if x.keys[i] == key:
                # remove key
                if x.is_leaf:
                    self.__remove_key_from_leaf(x, key)
                else:
                    self.__remove_key_from_node(x, key, i)
            else:
                if x.children[i].n >= self.min_degree:
                    self.__remove_internal(x.children[i], key)
                else:
                    if i > 0 and x.children[i - 1].n >= self.min_degree:
                        self.__sink_key_to_left_node(x, key, i)
                    elif x.children[i + 1].n >= self.min_degree:
                        self.__sink_key_to_right_node(x, key, i)
                    elif i > 0 and x.children[i - 1].n == self.min_degree - 1:
                        self.__merge_key_to_left_node(x, key, i)
                    else:
                        self.__merge_key_to_right_node(x, key, i)

    def __remove_key_from_leaf(self, x, key):
        x.keys.remove(key)
        x.n -= 1

    def __remove_key_from_node(self, x, key, i):
        if x.children[i].n >= self.min_degree:
            y,key_to_up = self.__find_right_before_key(x, i)
            key_to_down = x.keys[i]

            x.keys[i] = key_to_up
            y.keys = y.keys[:-1]
            y.keys.append(key_to_down)

            self.__remove_internal(x.children[i], key)
        elif x.children[i + 1].n >= self.min_degree:
            y, key_to_up = self.__find_right_after_key(x, i)
            key_to_down = x.keys[i]

            x.keys[i] = key_to_up
            y.keys = y.keys[1:]
            y.keys.insert(0, key_to_down)

            self.__remove_internal(x.children[i + 1], key)
        else:
            y = x.children[i]
            z = x.children[i + 1]
            key_to_down = x.keys[i]

            y.keys.append(key_to_down)
            y.n += 1
            y.keys.extend(z.keys)
            y.n += z.n
            x.children.remove(z)

            z.keys.clear()
            z.n = 0
            if not y.is_leaf:
                y.children.extend(z.children)
            z.children.clear()

            x.keys.remove(key_to_down)
            x.n -= 1
            if x.n == 0:
                self.root = y

            self.__remove_internal(y, key)

    def __find_right_before_key(self, x, i):
        y = x.children[i]

        while not y.is_leaf:
            y = y.children[-1]

        return y, y.keys[-1]

    def __find_right_after_key(self, x, i):
        y = x.children[i + 1]

        while not y.is_leaf:
            y = y.children[0]

        return y, y.keys[0]

    def __sink_key_to_left_node(self, x, key, i):
        y = x.children[i]
        z = x.children[i - 1]
        key_to_down = x.keys[i - 1]
        key_to_up = z.keys[-1]

        z.keys = z.keys[:-1]
        z.n -= 1
        x.keys[i - 1] = key_to_up
        y.keys.insert(0, key_to_down)
        y.n += 1

        if not y.is_leaf:
            children_to_move = z.children[-1]
            y.children.insert(0, children_to_move)
            z.children = z.children[:-1]

        self.__remove_internal(y, key)

    def __merge_key_to_left_node(self, x, key, i):
        y = x.children[i]
        z = x.children[i - 1]
        key_to_down = x.keys[i - 1]

        z.keys.append(key_to_down)
        z.n += 1
        z.keys.extend(y.keys)
        z.n += y.n
        if not z.is_leaf:
            z.children.extend(y.children)

        y.keys.clear()
        y.n = 0
        y.children.clear()

        x.keys.remove(key_to_down)
        x.n -= 1
        x.children.remove(y)

        if x.n == 0:
            self.root = z

        self.__remove_internal(z, key)

    def __merge_key_to_right_node(self, x, key, i):
        y = x.children[i]
        z = x.children[i + 1]
        key_to_down = x.keys[i]

        y.keys.append(key_to_down)
        y.n += 1
        y.keys.extend(z.keys)
        y.n += z.n
        if not y.is_leaf:
            y.children.extend(z.children)

        z.keys.clear()
        z.n = 0
        z.children.clear()

        x.keys.remove(key_to_down)
        x.n -= 1
        x.children.remove(z)

        if x.n == 0:
            self.root = y

        self.__remove_internal(y, key)

    def __sink_key_to_right_node(self, x, key, i):
        y = x.children[i]
        z = x.children[i + 1]
        key_to_down = x.keys[i]
        key_to_up = z.keys[0]

        z.keys = z.keys[1:]
        z.n -= 1
        x.keys[i] = key_to_up
        y.keys.append(key_to_down)
        y.n += 1

        if not y.is_leaf:
            children_to_move = z.children[0]
            y.children.append(children_to_move)
            z.children = z.children[1:]

        self.__remove_internal(y, key)


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
            if not has_key:
                print(f"{key} is not found in the tree.")
            else:
                t.remove(key)
        elif op == 3:
            t.print_tree()
        else:
            print("invalid operation")
