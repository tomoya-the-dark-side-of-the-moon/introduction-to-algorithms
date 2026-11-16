package main

import (
	"fmt"
	"math"
)

func cutRod(prices []int, length int) int {
	results := make([]int, length+1)
	for i := range length + 1 {
		results[i] = -math.MaxInt64
	}

	return cutRodInternal(prices, length, results)
}

func cutRodInternal(prices []int, length int, results []int) int {
	if results[length] > -math.MaxInt64 {
		return results[length]
	}

	var result int
	if length == 0 {
		result = 0
	} else {
		result = -math.MaxInt64
		for i := 1; i < length+1; i++ {
			result = max(result, prices[i]+cutRodInternal(prices, length-i, results))
		}
	}

	results[length] = result

	return result
}

func main() {
	prices := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

	fmt.Print("  i |")
	for i := range prices {
		fmt.Printf("%3d|", i)
	}
	fmt.Println()
	fmt.Print("p[i]|")
	for _, p := range prices {
		fmt.Printf("%3d|", p)
	}
	fmt.Println()

	for i := range prices {
		result := cutRod(prices, i)

		fmt.Printf("length = %d / result = %d\n", i, result)
	}
}
