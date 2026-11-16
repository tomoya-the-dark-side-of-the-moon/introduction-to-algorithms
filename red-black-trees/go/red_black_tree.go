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

func nilElement() *Element {
	if sentinel == nil {
		sentinel = &Element{key: -1, color: black}
	}

	return sentinel
}

func (t *Tree) put(z *Element) {
	t.size++

	var y *Element = nilElement()
	for x := t.root; x != nilElement(); {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.left = nilElement()
	z.right = nilElement()
	z.parent = y
	if y == nilElement() {
		t.root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}

	t.fixAfterInsertion(z)
}

func (t *Tree) contains(key int) *Element {
	for p := t.root; p != nilElement(); {
		if p.key == key {
			return p
		} else if key < p.key {
			p = p.left
		} else {
			p = p.right
		}
	}

	return nilElement()
}

func (t *Tree) successor(x *Element) *Element {
	if x.right != nilElement() {
		var p *Element
		for p = x.right; p.left != nilElement(); {
			p = p.left
		}

		return p
	} else {
		p := x
		q := p.parent
		for q != nilElement() && q.right == p {
			p = q
			q = p.parent
		}

		return q
	}
}

func (t *Tree) remove(z *Element) {
	t.size--

	y := z
	colorOfY := y.color
	var x *Element
	if z.left == nilElement() {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nilElement() {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.successor(z)
		x = y.right
		colorOfY = y.color
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		} else {
			x.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if colorOfY == black {
		t.fixAfterDeletion(x)
	}

}

func (t *Tree) printTree() {
	fmt.Printf("size = %d, height = %d\n", t.size, t.getHeight())

	if t.root != nilElement() {
		printTreeInternal(t.root, 0)
	}
}

func (t *Tree) transplant(x, y *Element) {
	if x.parent == nilElement() {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.parent = x.parent
}

func printTreeInternal(r *Element, spaceSize int) {
	if r.left != nilElement() {
		printTreeInternal(r.left, spaceSize+1)
	}

	for range spaceSize {
		fmt.Print("    ")
	}
	fmt.Printf("%d(", r.key)
	if r.color == black {
		fmt.Println("BLACK)")
	} else {
		fmt.Println("RED)")
	}

	if r.right != nilElement() {
		printTreeInternal(r.right, spaceSize+1)
	}
}

func (t *Tree) getHeight() int {
	if t.root == nilElement() {
		return 0
	}

	return t.getHeightInternal(t.root)
}

func (t *Tree) getHeightInternal(r *Element) int {
	if r == nilElement() {
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
	if y.left != nilElement() {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == nilElement() {
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
	if y.right != nilElement() {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == nilElement() {
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

func (t *Tree) fixAfterDeletion(x *Element) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				// case 1
				x.parent.color = red
				w.color = black
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				// case 2
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					// case 3
					w.color = red
					w.left.color = black
					t.rotateRight(w)
					w = x.parent.right
				}
				// case 4
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				// case 1
				x.parent.color = red
				w.color = black
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				// case 2
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					// case 3
					w.color = red
					w.right.color = black
					t.rotateLeft(w)
					w = x.parent.left
				}
				// case 4
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}

	x.color = black
}

func main() {
	t := &Tree{size: 0, height: 0, root: nilElement()}

	t.put(&Element{key: 2})
	t.put(&Element{key: 1})
	t.put(&Element{key: 4})
	t.put(&Element{key: 5})
	t.put(&Element{key: 9})
	t.put(&Element{key: 3})
	t.put(&Element{key: 6})
	t.put(&Element{key: 7})

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
			if z == nilElement() {
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
			if z == nilElement() {
				fmt.Printf("%d is not found in the tree.\n", key)
			} else {
				succ := t.successor(z)
				if succ == nilElement() {
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
