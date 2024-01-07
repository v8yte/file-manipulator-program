package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse [string to inputpath] [string to outputpath]",
	Short: `<inputpath> にあるファイルを受け取り、<outputpath> に <inputpath> の内容を逆にした新しいファイルを作成します。`,
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

        // データを逆にする
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}   

        err = os.WriteFile(outputfile, data, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
			os.Exit(1)
		}
    },
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
