package main

import (
	"fmt"
	"math"
)

func matrixChainOrder(p []int, idxOfStart, idxOfEnd int) int {
	results := make([][]int, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		results[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(p)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			results[i][j] = math.MaxInt64
		}
	}

	return matrixChainOrderInternal(p, idxOfStart, idxOfEnd, results)
}

func matrixChainOrderInternal(p []int, idxOfStart, idxOfEnd int, results [][]int) int {
	if results[idxOfStart][idxOfEnd] < math.MaxInt64 {
		return results[idxOfStart][idxOfEnd]
	}

	result := math.MaxInt64
	if idxOfStart == idxOfEnd {
		result = 0
	} else {
		for k := idxOfStart; k < idxOfEnd; k++ {
			result = min(result, matrixChainOrderInternal(p, idxOfStart, k, results)+matrixChainOrderInternal(p, k+1, idxOfEnd, results)+p[idxOfStart]*p[k+1]*p[idxOfEnd+1])
		}
	}

	results[idxOfStart][idxOfEnd] = result

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
