package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// copyCmd represents the reverse command
var copyCmd = &cobra.Command{
	Use:   "copy [string to inputpath] [string to outputpath]",
	Short: `<inputpath> にあるファイルのコピーを作成し、<outputpath> として保存します。`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1),cobra.MaximumNArgs(2)),
    Run: func(cmd *cobra.Command, args []string) {
        inputfile := args[0]
        outputfile := "output.txt"
        if len(args) == 2 {
            outputfile = args[1]
        }

        data, err := os.ReadFile(inputfile)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

        err = os.WriteFile(outputfile, data, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
			os.Exit(1)
		}
    },
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
