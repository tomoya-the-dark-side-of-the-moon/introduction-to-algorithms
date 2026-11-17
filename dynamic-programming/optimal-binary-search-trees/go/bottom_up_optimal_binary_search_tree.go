package main

import (
	"fmt"
	"math"
)

func optimalBst(p, q []float64) ([][]float64, [][]float64, [][]int) {
	results := make([][]float64, len(p)+1)
	w := make([][]float64, len(p)+1)
	roots := make([][]int, len(p)+1)

	for i := range results {
		results[i] = make([]float64, len(p)+1)
		w[i] = make([]float64, len(p)+1)
		roots[i] = make([]int, len(p)+1)
	}

	for i := 1; i < len(p)+1; i++ {
		results[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}

	for diff := 0; diff < len(p)-1; diff++ {
		i := 1
		j := i + diff

		for j < len(p) {
			w[i][j] = w[i][j-1] + p[j] + q[j]

			result := math.MaxFloat64
			for r := i; r < j+1; r++ {
				if result > results[i][r-1]+results[r+1][j]+w[i][j] {
					result = results[i][r-1] + results[r+1][j] + w[i][j]
					roots[i][j] = r
				}
			}

			results[i][j] = result

			i++
			j++
		}
	}

	return results, w, roots
}

func main() {
	p := []float64{0.0, 0.15, 0.10, 0.05, 0.10, 0.20}
	q := []float64{0.05, 0.10, 0.05, 0.05, 0.05, 0.10}

	fmt.Print("  i |")
	for i := range p {
		fmt.Printf("%4d|", i)
	}
	fmt.Println()
	fmt.Print("p[i]|")
	for _, _p := range p {
		fmt.Printf("%.2f|", _p)
	}
	fmt.Println()
	fmt.Print("q[i]|")
	for _, _q := range q {
		fmt.Printf("%.2f|", _q)
	}
	fmt.Println()

	results, w, roots := optimalBst(p, q)

	for diff := -1; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		i := 1
		j := i + diff
		for j < len(p) {

			fmt.Printf("(i, j) = (%d, %d) / result = %.2f / w = %.2f / root = %d\n", i, j, results[i][j], w[i][j], roots[i][j])

			i++
			j++
		}
	}
}
