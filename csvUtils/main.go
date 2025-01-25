package csvUtils

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MergeCsvFiles(cmd *cobra.Command, files []string) {
	fmt.Printf("Merging CSV files\n")
	fmt.Printf("Files: %v\n", files)
}
