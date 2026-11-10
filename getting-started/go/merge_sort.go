package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	mergeSort(data, 0, len(data)-1)
}

func mergeSort(data []int, idxOfStart, idxOfEnd int) {
	if idxOfStart < idxOfEnd {
		idxOfMid := (idxOfStart + idxOfEnd) / 2

		mergeSort(data, idxOfStart, idxOfMid)
		mergeSort(data, idxOfMid+1, idxOfEnd)

		merge(data, idxOfStart, idxOfMid, idxOfEnd)
	}
}

func merge(data []int, idxOfStart, idxOfMid, idxOfEnd int) {
	firstHalf := make([]int, idxOfMid-idxOfStart+1)
	secondHalf := make([]int, idxOfEnd-idxOfMid)
	copy(firstHalf, data[idxOfStart:idxOfMid+1])
	copy(secondHalf, data[idxOfMid+1:idxOfEnd+1])

	i, j := 0, 0
	for k := idxOfStart; k < idxOfEnd+1; k++ {
		if i < len(firstHalf) && j < len(secondHalf) {
			if firstHalf[i] < secondHalf[j] {
				data[k] = firstHalf[i]
				i++
			} else {
				data[k] = secondHalf[j]
				j++
			}
		} else if i < len(firstHalf) {
			data[k] = firstHalf[i]
			i++
		} else {
			data[k] = secondHalf[j]
			j++
		}
	}

	fmt.Println(data)
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
