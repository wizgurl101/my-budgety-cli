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

	var mergeCsvFiles = &cobra.Command{
		Use:   "merge-csv",
		Short: "Merge 2 CSV files with no duplicates",
		Run: func(cmd *cobra.Command, args []string) {
			csvUtils.MergeCsvFiles()
		},
	}

	rootCmd.AddCommand(mergeCsvFiles)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
