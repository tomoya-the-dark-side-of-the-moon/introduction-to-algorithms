package main

import "fmt"

func lcsLength(x, y string, i, j int) int {
	results := make([][]int, len(x)+1)
	for i := range len(x) + 1 {
		results[i] = make([]int, len(y)+1)
	}

	for i := range len(x) + 1 {
		for j := range len(y) + 1 {
			results[i][j] = -1
		}
	}

	return lcsLengthInternal(x, y, i, j, results)
}

func lcsLengthInternal(x, y string, i, j int, results [][]int) int {
	if results[i][j] >= 0 {
		return results[i][j]
	}

	var result int
	if i == 0 && j == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = 0
		}
	} else if i == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = lcsLengthInternal(x, y, i, j-1, results)
		}
	} else if j == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = lcsLengthInternal(x, y, i-1, j, results)
		}
	} else {
		if x[i] == y[j] {
			result = lcsLengthInternal(x, y, i-1, j-1, results) + 1
		} else {
			if lcsLengthInternal(x, y, i, j-1, results) > lcsLengthInternal(x, y, i-1, j, results) {
				result = lcsLengthInternal(x, y, i, j-1, results)
			} else {
				result = lcsLengthInternal(x, y, i-1, j, results)
			}
		}
	}

	results[i][j] = result

	return result
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"

	for i := range x {
		fmt.Printf("--- i = %d ---\n", i)

		for j := range y {
			result := lcsLength(x, y, i, j)

			fmt.Printf("(i, j) = (%d, %d) / x = %s, y = %s / rsult = %d\n", i, j, x[:i+1], y[:j+1], result)
		}
	}
}
