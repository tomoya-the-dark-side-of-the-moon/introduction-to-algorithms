package main

import (
	"fmt"
	"math/rand/v2"
)

func sort(data []int) {
	n := len(data)
	for j := 1; j < n; j++ {
		key := data[j]
		i := j - 1
		for i >= 0 && data[i] > key {
			data[i+1] = data[i]
			i--
		}

		data[i+1] = key

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
