package main

import "fmt"

type Color int

const (
	red Color = iota
	black
)

type Element struct {
	key                 int
	left, right, parent *Element
	color               Color
}

type Tree struct {
	size, height int
	root         *Element
}

var sentinel *Element

func nil_element() *Element {
	if sentinel == nil {
		sentinel = &Element{key: -1, color: black}
	}

	return sentinel
}

func (t *Tree) put(z *Element) {
	t.size++

	var y *Element = nil_element()
	for x := t.root; x != nil_element(); {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.left = nil_element()
	z.right = nil_element()
	z.parent = y
	if y == nil_element() {
		t.root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}

	t.fixAfterInsertion(z)

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

	if t.root != nil_element() {
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
	if r.left != nil_element() {
		printTreeInternal(r.left, space_size+1)
	}

	for range space_size {
		fmt.Print("    ")
	}
	fmt.Printf("%d(", r.key)
	if r.color == black {
		fmt.Println("BLACK)")
	} else {
		fmt.Println("RED)")
	}

	if r.right != nil_element() {
		printTreeInternal(r.right, space_size+1)
	}
}

func (t *Tree) getHeight() int {
	if t.root == nil_element() {
		return 0
	}

	return t.getHeightInternal(t.root)
}

func (t *Tree) getHeightInternal(r *Element) int {
	if r == nil_element() {
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

func (t *Tree) rotateLeft(x *Element) {
	y := x.right
	x.right = y.left
	if y.left != nil_element() {
		y.left.parent = y
	}

	y.parent = x.parent
	if x.parent == nil_element() {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *Tree) rotateRight(x *Element) {
	y := x.left
	x.left = y.right
	if y.right != nil_element() {
		y.right.parent = y
	}

	y.parent = x.parent
	if x.parent == nil_element() {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (t *Tree) fixAfterInsertion(z *Element) {
	for z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == red {
				// case 1
				z.parent.color = black
				z.parent.parent.color = red
				y.color = black
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// case 2
					z = z.parent
					t.rotateLeft(z)
				}
				// case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == red {
				// case 1
				z.parent.color = black
				z.parent.parent.color = red
				y.color = black
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					// case 2
					z = z.parent
					t.rotateRight(z)
				}
				// case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateLeft(z.parent.parent)
			}
		}
	}

	t.root.color = black
}

func main() {
	t := &Tree{size: 0, height: 0, root: nil_element()}

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
