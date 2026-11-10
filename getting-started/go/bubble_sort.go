package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if data[j+1] < data[j] {
				data[j+1], data[j] = data[j], data[j+1]
				swapped = true
			}
		}

		if !swapped {
			break
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
