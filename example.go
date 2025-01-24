package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Example() {
	var inputFiles []string
	var outputFile string

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "csvtool",
		Short: "A CLI tool for CSV operations",
		Long:  "csvtool is a command-line tool for performing various operations on CSV files, including merging without duplicates.",
	}

	// Merge CSV command
	var mergeCmd = &cobra.Command{
		Use:   "merge-csv",
		Short: "Merge CSV files with no duplicates",
		Long:  "Merge multiple CSV files and ensure there are no duplicate rows in the output.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(inputFiles) == 0 {
				fmt.Println("Please provide input CSV files using the --input flag.")
				return
			}

			if outputFile == "" {
				fmt.Println("Please provide an output file using the --output flag.")
				return
			}

			err := mergeCSVFiles(inputFiles, outputFile)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("CSV files merged successfully into %s\n", outputFile)
			}
		},
	}

	// Add flags to merge command
	mergeCmd.Flags().StringSliceVarP(&inputFiles, "input", "i", []string{}, "Input CSV files to merge")
	mergeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output CSV file")

	// Add the merge command to the root
	rootCmd.AddCommand(mergeCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mergeCSVFiles(inputFiles []string, outputFile string) error {
	recordsMap := make(map[string]bool)
	var mergedRecords [][]string

	for _, file := range inputFiles {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", file, err)
		}
		defer f.Close()

		reader := csv.NewReader(f)
		records, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read CSV from file %s: %v", file, err)
		}

		for _, record := range records {
			key := strings.Join(record, ",")
			if !recordsMap[key] {
				recordsMap[key] = true
				mergedRecords = append(mergedRecords, record)
			}
		}
	}

	// Write merged records to the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %v", outputFile, err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	err = writer.WriteAll(mergedRecords)
	if err != nil {
		return fmt.Errorf("failed to write to output file %s: %v", outputFile, err)
	}

	return nil
}
