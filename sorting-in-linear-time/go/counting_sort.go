package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func sort(data []int) {
	countingTable := make([]int, slices.Max(data)+1)

	for _, d := range data {
		countingTable[d]++
	}

	for i := 0; i < len(countingTable)-1; i++ {
		countingTable[i+1] += countingTable[i]
	}

	fmt.Println("--- counting table ---")
	fmt.Println(countingTable)

	temp := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		idxOfTarget := countingTable[data[i]] - 1
		temp[idxOfTarget] = data[i]
		countingTable[data[i]]--
	}

	data = data[:0]
	data = append(data, temp...)
}

func main() {
	n := 50
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
