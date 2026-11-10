package main

import (
	"fmt"
)

func search(data []int, key int) int {
	idxOfLow := 0
	idxOfHigh := len(data) - 1
	result := -1
	for idxOfLow <= idxOfHigh {
		idxOfMid := (idxOfLow + idxOfHigh) / 2
		if data[idxOfMid] == key {
			result = idxOfMid
			break
		} else if data[idxOfMid] < key {
			idxOfLow = idxOfMid + 1
		} else {
			idxOfHigh = idxOfMid - 1
		}
	}

	return result
}

func main() {
	n := 20
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i * i
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
