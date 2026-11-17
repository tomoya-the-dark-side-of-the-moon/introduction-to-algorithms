package main

import (
	"fmt"
	"math"
)

func matrixChainOrder(p []int) ([][]int, [][]int) {
	results := make([][]int, len(p)+1)
	thresholds := make([][]int, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		results[i] = make([]int, len(p)+1)
		thresholds[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(p)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			if i == j {
				results[i][j] = 0
				thresholds[i][j] = i
			} else {
				results[i][j] = math.MaxInt64
			}
		}
	}

	for diff := 1; diff < len(p)-1; diff++ {
		i := 0
		for j := diff; j < len(p)-1; j++ {
			result := math.MaxInt64
			for k := i; k < j; k++ {
				if result > results[i][k]+results[k+1][j]+p[i]*p[k+1]*p[j+1] {
					result = results[i][k] + results[k+1][j] + p[i]*p[k+1]*p[j+1]
					thresholds[i][j] = k
				}
			}

			results[i][j] = result

			i++
		}
	}

	return results, thresholds
}

func printOptimalStrategy(thresholds [][]int, i, j int) {
	if i == j {
		fmt.Printf("A%d", thresholds[i][j])
	} else {
		fmt.Print("(")
		printOptimalStrategy(thresholds, i, thresholds[i][j])
		printOptimalStrategy(thresholds, thresholds[i][j]+1, j)
		fmt.Print(")")
	}
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

	results, thresholds := matrixChainOrder(p)

	for diff := 0; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		i := 0
		for j := diff; j < len(p)-1; j++ {
			fmt.Printf("(i, j) = (%d, %d) / result = %d ", i, j, results[i][j])
			if i != j {
				printOptimalStrategy(thresholds, i, j)
			}
			fmt.Println()

			i++
		}
	}
}
