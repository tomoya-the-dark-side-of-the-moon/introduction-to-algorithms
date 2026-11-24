package main

import (
	"fmt"
	"slices"
)

type Element struct {
	keys     []int
	n        int
	children []*Element
	isLeaf   bool
}

type Tree struct {
	root      *Element
	minDegree int
}

func createElement() *Element {
	e := &Element{}
	e.keys = make([]int, 0)
	e.n = 0
	e.children = make([]*Element, 0)
	e.isLeaf = true

	return e
}

func (t *Tree) insert(key int) {
	r := t.root
	if r.n == 2*t.minDegree-1 {
		s := createElement()
		s.isLeaf = false
		s.children = slices.Insert(s.children, 0, r)

		t.root = s
		t.splitChild(s, 0)

		t.insertNonFull(s, key)
	} else {
		t.insertNonFull(r, key)
	}
}

func (t *Tree) insertNonFull(r *Element, key int) {
	if r.isLeaf {
		i := r.n - 1
		for i >= 0 && r.keys[i] > key {
			i--
		}

		i++

		r.keys = slices.Insert(r.keys, i, key)
		r.n++
	} else {
		i := r.n - 1
		for i >= 0 && r.keys[i] > key {
			i--
		}

		i++

		if r.children[i].n == 2*t.minDegree-1 {
			t.splitChild(r, i)

			if key > r.keys[i] {
				i++
			}
		}

		t.insertNonFull(r.children[i], key)
	}
}

func (t *Tree) splitChild(x *Element, i int) {
	y := x.children[i]
	z := createElement()

	keyToUp := y.keys[t.minDegree-1]

	z.keys = slices.Clone(y.keys[t.minDegree:])
	z.n = len(z.keys)
	y.keys = slices.Clone(y.keys[:t.minDegree-1])
	y.n = len(y.keys)

	z.isLeaf = y.isLeaf
	if !y.isLeaf {
		z.children = slices.Clone(y.children[t.minDegree:])
		y.children = slices.Clone(y.children[:t.minDegree])
	}

	x.keys = slices.Insert(x.keys, i, keyToUp)
	x.n++
	x.children = slices.Insert(x.children, i+1, z)
}

func (t *Tree) contains(key int) bool {
	return t.containsInternal(t.root, key)
}

func (t *Tree) containsInternal(x *Element, key int) bool {
	i := 0
	for i < x.n && x.keys[i] < key {
		i++
	}

	if i < x.n && x.keys[i] == key {
		return true
	} else if x.isLeaf {
		return false
	} else {
		return t.containsInternal(x.children[i], key)
	}
}

func (t *Tree) remove(key int) {
	t.removeInternal(t.root, key)
}

func (t *Tree) removeInternal(x *Element, key int) {
	i := 0
	for i < x.n && x.keys[i] < key {
		i++
	}

	if i == x.n {
		if x.children[i].n >= t.minDegree {
			t.removeInternal(x.children[i], key)
		} else {
			if x.children[i-1].n >= t.minDegree {
				t.liftKeyFromLeftNode(x, key, i)
			} else {
				t.mergeKeyToLeftNode(x, key, i)
			}
		}
	} else {
		if x.keys[i] == key {
			if x.isLeaf {
				t.removeKeyFromLeaf(x, key, i)
			} else {
				t.removeKeyFromNode(x, key, i)
			}
		} else {
			if x.children[i].n >= t.minDegree {
				t.removeInternal(x.children[i], key)
			} else if i > 0 && x.children[i-1].n >= t.minDegree {
				t.liftKeyFromLeftNode(x, key, i)
			} else if x.children[i+1].n >= t.minDegree {
				t.liftKeyFromRightNode(x, key, i)
			} else if i > 0 && x.children[i-1].n == t.minDegree-1 {
				t.mergeKeyToLeftNode(x, key, i)
			} else {
				t.mergeKeyToRightNode(x, key, i)
			}
		}
	}
}

func (t *Tree) removeKeyFromLeaf(x *Element, key, i int) {
	// case 1
	x.keys = slices.Delete(x.keys, i, i+1)
	x.n--
}

func (t *Tree) removeKeyFromNode(x *Element, key, i int) {
	if x.children[i].n >= t.minDegree {
		// case 2a
		y, keyToUp := t.findRightBeforeKey(x, i)
		keyToDown := x.keys[i]

		x.keys[i] = keyToUp
		y.keys = y.keys[:y.n-1]
		y.keys = append(y.keys, keyToDown)

		t.removeInternal(x.children[i], key)
	} else if x.children[i+1].n >= t.minDegree {
		// case 2b
		y, keyToUp := t.findRightAfterKey(x, i)
		keyToDown := x.keys[i]

		x.keys[i] = keyToUp
		y.keys = slices.Clone(y.keys[1:])
		y.keys = slices.Insert(y.keys, 0, keyToDown)

		t.removeInternal(x.children[i+1], key)
	} else {
		// case 2c
		y := x.children[i]
		z := x.children[i+1]
		keyToMerge := x.keys[i]

		y.keys = append(y.keys, keyToMerge)
		y.n += 1
		y.keys = append(y.keys, z.keys...)
		y.n += z.n

		if !y.isLeaf {
			y.children = append(y.children, z.children...)
		}

		x.keys = slices.Delete(x.keys, i, i+1)
		x.n--
		x.children = slices.Delete(x.children, i+1, i+2)

		clear(z.keys)
		z.keys = z.keys[:0]
		z.n = 0
		clear(z.children)
		z.children = z.children[:0]

		if x.n == 0 {
			t.root = y
		}

		t.removeInternal(y, key)
	}
}

func (t *Tree) findRightBeforeKey(x *Element, i int) (*Element, int) {
	y := x.children[i]

	for !y.isLeaf {
		y = y.children[y.n]
	}

	return y, y.keys[y.n-1]
}

func (t *Tree) findRightAfterKey(x *Element, i int) (*Element, int) {
	y := x.children[i+1]

	for !y.isLeaf {
		y = y.children[0]
	}

	return y, y.keys[0]
}

func (t *Tree) liftKeyFromLeftNode(x *Element, key, i int) {
	// case 3a-1
	y := x.children[i]
	z := x.children[i-1]
	keyToDown := x.keys[i-1]
	keyToUp := z.keys[z.n-1]

	z.keys = slices.Clone(z.keys[:z.n-1])
	z.n--
	x.keys[i-1] = keyToUp
	y.keys = slices.Insert(y.keys, 0, keyToDown)
	y.n++

	if !z.isLeaf {
		childrenToMove := z.children[len(z.children)-1]
		y.children = slices.Insert(y.children, 0, childrenToMove)
		z.children = slices.Clone(z.children[:len(z.children)-1])
	}

	t.removeInternal(y, key)
}

func (t *Tree) liftKeyFromRightNode(x *Element, key, i int) {
	// case 3a-2
	y := x.children[i]
	z := x.children[i+1]
	keyToDown := x.keys[i]
	keyToUp := z.keys[0]

	z.keys = slices.Clone(z.keys[1:])
	z.n--
	x.keys[i] = keyToUp
	y.keys = append(y.keys, keyToDown)
	y.n++

	if !z.isLeaf {
		childrenToMove := z.children[0]
		y.children = append(y.children, childrenToMove)
		z.children = slices.Clone(z.children[1:])
	}

	t.removeInternal(y, key)
}

func (t *Tree) mergeKeyToLeftNode(x *Element, key, i int) {
	// case 3b-1
	y := x.children[i]
	z := x.children[i-1]
	keyToDown := x.keys[i-1]

	z.keys = append(z.keys, keyToDown)
	z.n++
	z.keys = append(z.keys, y.keys...)
	z.n += y.n
	if !z.isLeaf {
		z.children = append(z.children, y.children...)
	}

	x.keys = slices.Delete(x.keys, i-1, i)
	x.n--
	x.children = slices.Delete(x.children, i, i+1)
	if x.n == 0 {
		t.root = z
	}

	clear(y.keys)
	y.keys = y.keys[:0]
	y.n = 0
	clear(y.children)
	y.children = y.children[:0]

	t.removeInternal(z, key)
}

func (t *Tree) mergeKeyToRightNode(x *Element, key, i int) {
	// case 3b-2
	y := x.children[i]
	z := x.children[i+1]
	keyToDown := x.keys[i]

	y.keys = append(y.keys, keyToDown)
	y.n++
	y.keys = append(y.keys, z.keys...)
	y.n += z.n
	if !y.isLeaf {
		y.children = append(y.children, z.children...)
	}

	x.keys = slices.Delete(x.keys, i, i+1)
	x.n--
	x.children = slices.Delete(x.children, i+1, i+2)
	if x.n == 0 {
		t.root = y
	}

	clear(z.keys)
	z.keys = z.keys[:0]
	z.n = 0
	clear(z.children)
	z.children = z.children[:0]

	t.removeInternal(y, key)
}

func (t *Tree) printTree() {
	t.printTreeInternal(t.root, 0)
}

func (t *Tree) printTreeInternal(r *Element, spaceSize int) {
	if r.n == 0 {
		return
	}

	for range spaceSize {
		fmt.Print("  ")
	}
	fmt.Println(r.keys)

	for _, child := range r.children {
		t.printTreeInternal(child, spaceSize+1)
	}
}

func main() {
	t := &Tree{root: createElement(), minDegree: 2}

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

	for {
		fmt.Print("1:insert 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			t.insert(key)
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			hasKey := t.contains(key)
			if !hasKey {
				fmt.Printf("%d is not found in the tree.\n", key)
			} else {
				t.remove(key)
			}
		case 3:
			t.printTree()
		default:
			fmt.Println("invalid operation")
		}
	}
}
