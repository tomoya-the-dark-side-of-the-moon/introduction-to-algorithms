package main

import (
	"fmt"
	"math"
)

func matrixChainOrder(p []int, idxOfStart, idxOfEnd int) int {
	if idxOfStart == idxOfEnd {
		return 0
	}

	result := math.MaxInt64
	for k := idxOfStart; k < idxOfEnd; k++ {
		result = min(result, matrixChainOrder(p, idxOfStart, k)+matrixChainOrder(p, k+1, idxOfEnd)+p[idxOfStart]*p[k+1]*p[idxOfEnd+1])
	}

	return result
}

func main() {
	p := []int{30, 35, 15, 5, 10, 20, 25}

	fmt.Print("  i |")
	for i := range len(p) - 1 {
		fmt.Printf("%7d|", i)
	}
	fmt.Println()
	fmt.Print("A[i]|")
	for i := range len(p) - 1 {
		fmt.Printf("%3d*%3d|", p[i], p[i+1])
	}
	fmt.Println()

	for diff := 0; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		i := 0
		for j := diff; j < len(p)-1; j++ {
			result := matrixChainOrder(p, i, j)

			fmt.Printf("(i, j) = (%d, %d) / result = %d\n", i, j, result)

			i++
		}
	}
}
