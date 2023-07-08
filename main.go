package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

// Returns the max value from a slice of ints
func maxInt(slice []int) int {
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// Returns the max value from a slice of float64s
func maxFloat64(slice []float64) float64 {
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// Returns the max value from a slice of any type that implements Ordered
func maxGeneric[T constraints.Ordered](slice []T) T {
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9}
	float64Slice := []float64{2.7, 1.4, 2.7, 1.8, 2.8, 1.6}

	startTime := time.Now()
	fmt.Println(maxGeneric[int](intSlice))         // Output: 9
	fmt.Println(maxGeneric[float64](float64Slice)) // Output: 2.8
	fmt.Println("Time to execute generic: ", time.Since(startTime))

	startTime = time.Now()
	fmt.Println(maxInt(intSlice))         // Output: 9
	fmt.Println(maxFloat64(float64Slice)) // Output: 2.8
	fmt.Println("Time to execute non-generic: ", time.Since(startTime))
}
