package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 64}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {

	min, max := 0, 0
	minI, maxI := 0, 0

	for i, a := range prices {

		if a < min {
			min = a
			minI = i
		}
		if a > max && i != 0 {
			max = a
			maxI = i
		}

	}

	if maxI > minI {
		return max - min
	} else {
		return 0
	}
}
