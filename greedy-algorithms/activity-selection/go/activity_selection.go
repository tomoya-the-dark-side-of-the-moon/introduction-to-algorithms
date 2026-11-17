package main

import "fmt"

func selectActivity(s, f []int) []int {
	results := make([]int, 0)

	results = append(results, 0)
	finished := f[0]

	for i := 1; i < len(s); i++ {
		if finished <= s[i] {
			results = append(results, i)

			finished = f[i]
		}
	}

	return results
}

func main() {
	s := []int{1, 3, 0, 5, 3, 5, 6, 7, 8, 2, 12}
	f := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}

	fmt.Print("  i |")
	for i := range s {
		fmt.Printf("%3d|", i)
	}
	fmt.Println()
	fmt.Print("s[i]|")
	for _, _s := range s {
		fmt.Printf("%3d|", _s)
	}
	fmt.Println()
	fmt.Print("f[i]|")
	for _, _f := range f {
		fmt.Printf("%3d|", _f)
	}
	fmt.Println()

	results := selectActivity(s, f)

	fmt.Println("--- result ---")
	for _, i := range results {
		fmt.Printf("activity = %d / start = %d, finish = %d\n", i, s[i], f[i])
	}
}
