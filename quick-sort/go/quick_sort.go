package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, idxOfStart, idxOfEnd int) {
	if idxOfStart < idxOfEnd {
		idxOfPartition := partition(data, idxOfStart, idxOfEnd)

		fmt.Println(data)

		quickSort(data, idxOfStart, idxOfPartition-1)
		quickSort(data, idxOfPartition+1, idxOfEnd)
	}
}

func partition(data []int, idxOfStart, idxOfEnd int) int {
	pivot := data[idxOfEnd]

	i := idxOfStart
	for j := idxOfStart; j < idxOfEnd; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	data[i], data[idxOfEnd] = data[idxOfEnd], data[i]

	return i
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
