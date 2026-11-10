package main

import (
	"fmt"
	"math/rand/v2"
)

func search(data []int, key int) int {
	result := -1
	for i, d := range data {
		if d == key {
			result = i
			break
		}
	}

	return result
}

func main() {
	n := 20
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.IntN(20)
	}

	fmt.Println("--- data ---")
	fmt.Println(data)

	var key int
	fmt.Print("input key > ")
	fmt.Scanf("%d", &key)

	result := search(data, key)
	if result != -1 {
		fmt.Printf("result = %d\n", result)
	} else {
		fmt.Printf("%d is not found.\n", key)
	}
}
