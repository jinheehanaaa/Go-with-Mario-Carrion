package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Pipeline: We output a channel that will be the input of the next step
func main() {
	recordsC, err := readCSV("file1.csv")
	if err != nil {
		log.Fatalf("Could not read csv %v", err)
	}

	// Change first colimn's first word as upper-case
	for val := range sanitize(titlize(recordsC)) {
		fmt.Printf("%v\n", val)
	}
}

// Read values
func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	ch := make(chan []string)

	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}

// Remove "invalid" records
// Get rid of anything longer than 3 character
func sanitize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strC {
			if len(val[0]) > 3 {
				fmt.Println("skipped ", val)
				continue
			}

			ch <- val
		}

		close(ch)
	}()

	return ch
}

// Modify received values
// Take 1st chracter of 1st column and make upper-case
func titlize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strC {
			val[0] = strings.Title(val[0])
			val[1], val[2] = val[2], val[1]

			ch <- val
		}

		close(ch)
	}()

	return ch
}
