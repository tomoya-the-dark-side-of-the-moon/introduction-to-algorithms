package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	ch          string
	freq        int
	left, right *Element
}

type ElementHeap []*Element

func (h ElementHeap) Len() int {
	return len(h)
}

func (h ElementHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h ElementHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ElementHeap) Push(x any) {
	item := x.(*Element)
	*h = append(*h, item)
}

func (h *ElementHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func huffmanCode(characters string, frequencies []int) *Element {
	elements := &ElementHeap{}
	heap.Init(elements)

	for i := range characters {
		element := &Element{ch: string(characters[i]), freq: frequencies[i]}
		heap.Push(elements, element)
	}

	for elements.Len() > 1 {
		x := heap.Pop(elements)
		y := heap.Pop(elements)
		z := &Element{ch: string(""), freq: x.(*Element).freq + y.(*Element).freq, left: x.(*Element), right: y.(*Element)}

		heap.Push(elements, z)
	}

	return heap.Pop(elements).(*Element)
}

func printTree(root *Element) {
	printTreeInternal(root, 0)
}

func printTreeInternal(r *Element, spaceSize int) {
	if r.left != nil {
		printTreeInternal(r.left, spaceSize+1)
	}

	for range spaceSize {
		fmt.Print("  ")
	}
	fmt.Printf("%s(%d)\n", r.ch, r.freq)

	if r.right != nil {
		printTreeInternal(r.right, spaceSize+1)
	}
}

func main() {
	characters := "abcdef"
	frequencies := []int{45, 13, 12, 16, 9, 5}

	fmt.Print("  i |")
	for i := range characters {
		fmt.Printf("%3d|", i)
	}
	fmt.Println()
	fmt.Print(" ch |")
	for _, ch := range characters {
		fmt.Printf("%3c|", ch)
	}
	fmt.Println()
	fmt.Print("freq|")
	for _, freq := range frequencies {
		fmt.Printf("%3d|", freq)
	}
	fmt.Println()

	root := huffmanCode(characters, frequencies)
	printTree(root)
}
