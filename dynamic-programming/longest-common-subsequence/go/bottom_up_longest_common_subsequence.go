package main

import "fmt"

type direction int

const (
	diagonal = iota
	up
	left
)

func lcsLength(x, y string) ([][]int, [][]int) {
	results := make([][]int, len(x)+1)
	directions := make([][]int, len(x)+1)
	for i := range len(x) + 1 {
		results[i] = make([]int, len(y)+1)
		directions[i] = make([]int, len(y)+1)
	}

	for i := range len(x) {
		var result int
		for j := range len(y) {
			if i == 0 && j == 0 {
				if x[i] == y[j] {
					result = 1
					directions[i][j] = diagonal
				} else {
					result = 0
					directions[i][j] = up
				}
			} else if i == 0 {
				if x[i] == y[j] {
					result = 1
					directions[i][j] = diagonal
				} else {
					result = results[i][j-1]
					directions[i][j] = left
				}
			} else if j == 0 {
				if x[i] == y[j] {
					result = 1
					directions[i][j] = diagonal
				} else {
					result = results[i-1][j]
					directions[i][j] = up
				}
			} else {
				if x[i] == y[j] {
					result = results[i-1][j-1] + 1
					directions[i][j] = diagonal
				} else {
					if results[i][j-1] > results[i-1][j] {
						result = results[i][j-1]
						directions[i][j] = left
					} else {
						result = results[i-1][j]
						directions[i][j] = up
					}
				}
			}

			results[i][j] = result
		}
	}

	return results, directions
}

func printOptimalStrategy(x string, i, j int, directions [][]int) {
	if i < 0 || j < 0 {
		return
	} else {
		if directions[i][j] == diagonal {
			printOptimalStrategy(x, i-1, j-1, directions)
			fmt.Printf("%s", x[i:i+1])
		} else if directions[i][j] == up {
			printOptimalStrategy(x, i-1, j, directions)
		} else {
			printOptimalStrategy(x, i, j-1, directions)
		}
	}
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"

	results, directions := lcsLength(x, y)
	for i := range x {
		fmt.Printf("--- i = %d ---\n", i)

		for j := range y {
			fmt.Printf("(i, j) = (%d, %d) / x = %s, y = %s / rsult = %d (", i, j, x[:i+1], y[:j+1], results[i][j])
			printOptimalStrategy(x, i, j, directions)
			fmt.Println(")")
		}
	}
}
