package main

import (
	"fmt"
	"math"
)

func cutRod(prices []int) ([]int, []int) {
	results := make([]int, len(prices)+1)
	divided := make([]int, len(prices)+1)
	for i := range results {
		results[i] = -math.MaxInt64
	}

	results[0] = 0
	divided[0] = 0

	for j := 1; j < len(prices); j++ {
		result := -math.MaxInt64
		for i := 0; i < j+1; i++ {
			if result < prices[i]+results[j-i] {
				result = prices[i] + results[j-i]
				divided[j] = i
			}
		}

		results[j] = result
	}

	return results, divided
}

func printOptimalStrategy(length int, divided []int) {
	fmt.Print("(")

	for rest := length; rest > 0; rest -= divided[rest] {
		fmt.Printf("%d", divided[rest])
		if rest-divided[rest] > 0 {
			fmt.Print(" + ")
		}
	}

	fmt.Print(")")
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

	results, divided := cutRod(prices)
	for i := range prices {
		fmt.Printf("length = %d / result = %d ", i, results[i])
		if i != 0 {
			printOptimalStrategy(i, divided)
		}
		fmt.Println()
	}
}
