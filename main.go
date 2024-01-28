package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Parse command line arguments
	op := flag.String("op", "sum", "Operation to be executed")
	flag.Parse()

	if err := run(flag.Args(), *op, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, out io.Writer) error {
	
	// Stores a calculation function according to user-provided parameter
	var opFunc statsFunc

	// Check if no files have been provided
	if len(filenames) == 0 {
		return ErrNoFiles
	}

	// Validate user-provided operation
	switch op {
	case "sum":
		opFunc = sum
	case "mean":
		opFunc = mean
	case "min":
		opFunc = min
	case "max":
		opFunc = max
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)

	// Loop thru the provided files and store data in consolidate
	for _, fname := range filenames {
		
		// Open the file
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("Cannot open file: %w", err)
		}

		// Parse the file and store the floats in a slice 
		data, err := csv2float(f)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}

	_, err := fmt.Fprintln(out, opFunc(consolidate))
	return err
}