package main

import "fmt"

func lcsLength(x, y string, i, j int) int {
	if i == 0 && j == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return 0
		}
	} else if i == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return lcsLength(x, y, i, j-1)
		}
	} else if j == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return lcsLength(x, y, i-1, j)
		}
	} else {
		if x[i] == y[j] {
			return lcsLength(x, y, i-1, j-1) + 1
		} else {
			if lcsLength(x, y, i, j-1) > lcsLength(x, y, i-1, j) {
				return lcsLength(x, y, i, j-1)
			} else {
				return lcsLength(x, y, i-1, j)
			}
		}
	}
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
