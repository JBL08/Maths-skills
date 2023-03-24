package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	//opens the file specified by the first command line argument.

	file, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(1)
		}
		values = append(values, value)
	}

	if err := scanner.Err(); err != nil { //checks for errors during the scan and exits if there is
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if len(values) == 0 { //exits if no values found.
		fmt.Println("No values found in file.")
		os.Exit(1)
	}

	sort.Ints(values) //sorting in ascend order

	// Calculates and prints average
	var sum int
	for _, value := range values {
		sum += value
	}

	average := float64(sum) / float64(len(values))
	fmt.Printf("Average: %d\n", int(average))

	// Calculate and print median
	var median float64
	if len(values)%2 == 0 {
		median = (float64(values[len(values)/2-1]) + float64(values[len(values)/2])) / 2
	} else {
		median = float64(values[len(values)/2])
	}
	fmt.Printf("Median: %d\n", int(median))

	// Calculate and print variance
	var variance float64
	for _, value := range values {
		variance += math.Pow(float64(value)-average, 2)
	}
	variance /= float64(len(values))
	fmt.Printf("Variance: %d\n", int(variance))

	// Calculate and print standard deviation
	standardDeviation := math.Sqrt(variance)
	fmt.Printf("Standard Deviation: %d\n", int(standardDeviation))
}
