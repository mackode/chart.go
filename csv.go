package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"time"
)

type Viewing struct {
	Title	   string
	Date	   time.Time
}

func readHistory() ([]Viewing, error) {
	csvFile := "history.csv"
	viewings := []Viewing{}
	r, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", csvFile, err)
		return viewings, err
	}
	defer r.Close()

	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV file %s: %v\n", csvFile, err)
		return viewings, err
	}

	records = records[1:] // Skip header
	for _, line := range records {
		t, err := time.Parse("1/2/06", line[1])
		if err != nil {
			fmt.Printf("Error parsing date %s: %v\n", line, err)
			return viewings, err
		}
		v := Viewing{Title: line[0], Date: t}
		viewings = append(viewings, v)
	}

	return viewings, nil
}