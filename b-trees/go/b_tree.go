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

	//z.keys = y.keys[t.minDegree:]
	z.keys = slices.Clone(y.keys[t.minDegree:])
	z.n = len(z.keys)
	//y.keys = y.keys[:t.minDegree-1]
	y.keys = slices.Clone(y.keys[:t.minDegree-1])
	y.n = len(y.keys)

	z.isLeaf = y.isLeaf
	if !y.isLeaf {
		//z.children = y.children[t.minDegree:]
		z.children = slices.Clone(y.children[t.minDegree:])
		//y.children = y.children[:t.minDegree]
		y.children = slices.Clone(y.children[:t.minDegree])
	}

	x.keys = slices.Insert(x.keys, i, keyToUp)
	x.n++
	x.children = slices.Insert(x.children, i+1, z)
}

func (t *Tree) printTree() {
	t.printTreeInternal(t.root, 0)
}

func (t *Tree) printTreeInternal(r *Element, spaceSize int) {
	if r == nil {
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
		case 3:
			t.printTree()
		default:
			fmt.Println("invalid operation")
		}
	}
}
