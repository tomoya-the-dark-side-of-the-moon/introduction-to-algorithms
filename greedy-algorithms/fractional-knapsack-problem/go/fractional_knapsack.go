package main

import (
	"fmt"
	"sort"
)

func selectItems(weights, prices []float64, limit float64) map[int]float64 {
	unitPrices := make(map[int]float64)

	for i := range prices {
		unitPrices[i] = prices[i] / weights[i]
	}

	keys := make([]int, 0, len(unitPrices))
	for k := range unitPrices {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return unitPrices[keys[i]] > unitPrices[keys[j]]
	})

	results := make(map[int]float64)
	rest := limit
	for _, k := range keys {
		if rest >= weights[k] {
			results[k] = float64(weights[k])
		} else {
			results[k] = rest
		}

		rest -= results[k]
	}

	return results
}

func main() {
	weights := []float64{20, 30, 10}
	prices := []float64{100, 120, 60}

	fmt.Print("  i |")
	for i := range weights {
		fmt.Printf("%5d|", i)
	}
	fmt.Println()
	fmt.Print("w[i]|")
	for _, w := range weights {
		fmt.Printf("%5.1f|", w)
	}
	fmt.Println()
	fmt.Print("p[i]|")
	for _, p := range prices {
		fmt.Printf("%5.1f|", p)
	}
	fmt.Println()

	results := selectItems(weights, prices, 50)

	fmt.Println("--- results ---")
	total := 0.0
	for idx, weight := range results {
		fmt.Printf("index = %d, weight = %.1f\n", idx, weight)

		total += (prices[idx] / weights[idx]) * weight
	}
	fmt.Printf("total = %.1f\n", total)
}
