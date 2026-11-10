package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idxOfMin := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[idxOfMin] {
				idxOfMin = j
			}
		}

		if idxOfMin != i {
			data[i], data[idxOfMin] = data[idxOfMin], data[i]
		}

		fmt.Println(data)
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
