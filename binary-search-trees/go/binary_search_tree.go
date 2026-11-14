package main

import "fmt"

type Element struct {
	key                 int
	left, right, parent *Element
}

type Tree struct {
	size, height int
	root         *Element
}

func (t *Tree) put(z *Element) {
	t.size++

	var y *Element
	for x := t.root; x != nil; {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == nil {
		t.root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}

	t.height = t.getHeight()
}

func (t *Tree) contains(key int) *Element {
	for p := t.root; p != nil; {
		if p.key == key {
			return p
		} else if key < p.key {
			p = p.left
		} else {
			p = p.right
		}
	}

	return nil
}

func (t *Tree) successor(x *Element) *Element {
	if x.right != nil {
		var p *Element
		for p = x.right; p.left != nil; {
			p = p.left
		}

		return p
	} else {
		p := x
		q := p.parent
		for q != nil && q.right == p {
			p = q
			q = p.parent
		}

		return q
	}
}

func (t *Tree) remove(z *Element) {
	t.size--

	if z.left == nil {
		t.transplant(z, z.right)
	} else if z.right == nil {
		t.transplant(z, z.left)
	} else {
		y := t.successor(z)
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}

	z = nil

	t.height = t.getHeight()
}

func (t *Tree) printTree() {
	fmt.Printf("size = %d, height = %d\n", t.size, t.height)

	if t.root != nil {
		printTreeInternal(t.root, 0)
	}
}

func (t *Tree) transplant(x, y *Element) {
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	if y != nil {
		y.parent = x.parent
	}
}

func printTreeInternal(r *Element, space_size int) {
	if r.left != nil {
		printTreeInternal(r.left, space_size+1)
	}

	for range space_size {
		fmt.Print("    ")
	}
	fmt.Println(r.key)

	if r.right != nil {
		printTreeInternal(r.right, space_size+1)
	}
}

func (t *Tree) getHeight() int {
	if t.root == nil {
		return 0
	}

	return t.getHeightInternal(t.root)
}

func (t *Tree) getHeightInternal(r *Element) int {
	if r == nil {
		return -1
	}

	heightOfLeft := t.getHeightInternal(r.left)
	heightOfRight := t.getHeightInternal(r.right)
	if heightOfLeft < heightOfRight {
		return heightOfRight + 1
	} else {
		return heightOfLeft + 1
	}
}

func main() {
	t := &Tree{size: 0, height: 0, root: nil}

	t.put(&Element{key: 8})
	t.put(&Element{key: 3})
	t.put(&Element{key: 10})
	t.put(&Element{key: 1})
	t.put(&Element{key: 6})
	t.put(&Element{key: 14})
	t.put(&Element{key: 4})
	t.put(&Element{key: 7})
	t.put(&Element{key: 13})

	for {
		fmt.Print("1:put 2:remove 3:print 4:successor > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			t.put(&Element{key: key})
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			z := t.contains(key)
			if z == nil {
				fmt.Printf("%d is not found in the tree.\n", key)
			} else {
				t.remove(z)
			}
		case 3:
			t.printTree()
		case 4:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			z := t.contains(key)
			if z == nil {
				fmt.Printf("%d is not found in the tree.\n", key)
			} else {
				succ := t.successor(z)
				if succ == nil {
					fmt.Println("successor is nil.")
				} else {
					fmt.Printf("successor = %d\n", succ.key)
				}
			}
		default:
			fmt.Println("invalid operation")
		}
	}
}
