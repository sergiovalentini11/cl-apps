package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"slices"
)

// Sums all the numbers in a given slice
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}

// Uses sum to calculate the mean of a given slice
func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// Returns the miminum value in a given slice of floats
func min(data []float64) float64 {
	return slices.Min(data)
}

// Returns the maximum value in a given slice of floats
func max(data []float64) float64 {
	return slices.Max(data)
}

// Use this type for generic statistical functions
type statsFunc func(data []float64) float64

// Parse a CSV file to create a slice of floats. It accepts the CSV source and the column number as arguments
func csv2float(r io.Reader) ([]float64, error) {

	// Reads in data from CSV files
	cr := csv.NewReader(r)

	// Read in the CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Cannot read data from file: %w", err)
	}

	// CSV data will be strings, creating a slice to store the data after it's converted to float
	var data []float64

	// Iterate thru the whole file
	for i, row := range allData {
		if i == 0 {
			continue
		}

		for col := 0; col < len(row); col++ {
		// Try to convert column into a float
		v, err := strconv.ParseFloat(row[col], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
		}
	}
	// Return the slice of floats and nil error
	return data, nil
}
