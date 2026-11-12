package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func sort(data []int) {
	bucket := make([][]int, slices.Max(data)/10+1)
	for i := range bucket {
		bucket[i] = make([]int, 0, len(data))
	}
	for _, d := range data {
		bucket[d/10] = append(bucket[d/10], d)
	}

	fmt.Println("--- bucket after initializing ---")
	for i, b := range bucket {
		fmt.Printf("%d: %v\n", i, b)
	}

	for _, b := range bucket {
		slices.Sort(b)
	}

	fmt.Println("--- bucket after sorting ---")
	for i, b := range bucket {
		fmt.Printf("%d: %v\n", i, b)
	}

	data = data[:0]
	for _, b := range bucket {
		data = append(data, b...)
	}
}

func main() {
	n := 100
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.IntN(101)
	}

	fmt.Println("--- before ---")
	fmt.Println(data)

	fmt.Println("--- sort ---")
	sort(data)

	fmt.Println("--- after ---")
	fmt.Println(data)
}
