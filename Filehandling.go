package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func copyCSVData() {
	// Open the source CSV file
	sourceFile, err := os.Open("source.csv")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Read the CSV data
	reader := csv.NewReader(sourceFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}

	// Open the destination CSV file
	destinationFile, err := os.Create("destination.csv")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	// Write the records to the destination file
	writer := csv.NewWriter(destinationFile)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing CSV data:", err)
		return
	}

	fmt.Println("CSV data copied successfully.")
}
