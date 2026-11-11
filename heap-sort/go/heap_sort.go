package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	buildMaxHeap(data)

	fmt.Println(data)

	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]

		fmt.Println(data)

		heapify(data, 0, i-1)
	}
}

func buildMaxHeap(data []int) {
	for i := len(data) / 2; i >= 0; i-- {
		heapify(data, i, len(data)-1)
	}
}

func heapify(data []int, idxOfTarget, idxOfHeap int) {
	idxOfLargest := idxOfTarget

	idxOfLeft := 2*idxOfTarget + 1
	if idxOfLeft <= idxOfHeap && data[idxOfLeft] > data[idxOfLargest] {
		idxOfLargest = idxOfLeft
	}
	idxOfRight := 2 * (idxOfTarget + 1)
	if idxOfRight <= idxOfHeap && data[idxOfRight] > data[idxOfLargest] {
		idxOfLargest = idxOfRight
	}

	if idxOfLargest != idxOfTarget {
		data[idxOfTarget], data[idxOfLargest] = data[idxOfLargest], data[idxOfTarget]

		heapify(data, idxOfLargest, idxOfHeap)
	}
}

func main() {
	n := 20
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.IntN(20)
	}

	fmt.Println("--- before ---")
	fmt.Println(data)

	fmt.Println("--- sort ---")
	sort(data)

	fmt.Println("--- after ---")
	fmt.Println(data)
}
