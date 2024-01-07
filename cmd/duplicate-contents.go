package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// duplicate-contentsCmd represents the duplicate-contents command
var duplicate_contentsCmd = &cobra.Command{
	Use:   "duplicate-contents [string to inputpath] [string to outputpath]",
	Short: `<inputpath> にあるファイルの内容を読み込み、その内容を複製し、複製された内容を <inputpath> に n 回複製します。`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(2),cobra.MaximumNArgs(2)),
    Run: func(cmd *cobra.Command, args []string) {
        inputfile := args[0]
        outputfile := "output.txt"
        n, err := strconv.Atoi(args[1])

        data, err := os.ReadFile(inputfile)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

        // n回複製する。
        var duplicatedData []byte
        for i := 0; i < n; i++ {
            duplicatedData = append(duplicatedData, data...)
		}   

        err = os.WriteFile(outputfile, duplicatedData, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
			os.Exit(1)
		}
    },
}

func init() {
	rootCmd.AddCommand(duplicate_contentsCmd)
}
