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

	var year int
	var startMonth int
	var setYearBudgetAmount = &cobra.Command{
		Use:   "set-year-budget",
		Short: "Set budget amount for a year",
		Run: func(cmd *cobra.Command, args []string) {
			if startMonth < 1 || startMonth > 12 {
				fmt.Printf("Invalid start month. Please provide a valid month between 1 and 12\n")
				return
			}

			if year < 1970 || year > 3000 {
				fmt.Printf("Invalid year. Please provide a valid year\n")
				return
			}

			fmt.Printf("Set Budget For Year")
		},
	}
	rootCmd.AddCommand(setYearBudgetAmount)
	setYearBudgetAmount.Flags().IntVarP(&year, "year", "y", 0, "Year for which budget amount needs to be set")
	setYearBudgetAmount.Flags().IntVarP(&startMonth, "start-month", "s", 1, "Start month for the budget")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
