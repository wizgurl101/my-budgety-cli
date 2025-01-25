package main

import (
	"fmt"
	csvUtils "my-budgety-cli/csvUtils"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mybudgetycli",
		Short: "My Budgety CLI",
		Long:  "My Budgety CLI is a command-line tool to manage your budget and expenses.",
	}

	var filePaths []string
	var mergeCsvFiles = &cobra.Command{
		Use:   "merge-csv",
		Short: "Merge 2 CSV files with no duplicates",
		Run: func(cmd *cobra.Command, args []string) {
			if len(filePaths) == 0 {
				fmt.Println("Please provide input CSV files using the --file flag.")
				return
			}

			csvUtils.MergeCsvFiles(cmd, filePaths)
		},
	}
	mergeCsvFiles.Flags().StringSliceVarP(&filePaths, "files", "f", []string{}, "Input CSV files to merge")

	rootCmd.AddCommand(mergeCsvFiles)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
