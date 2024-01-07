package cmd

import (
	"fmt"
	"os"
    "strings"

	"github.com/spf13/cobra"
)

// replace-stringCmd represents the replace-string command
var replace_stringCmd = &cobra.Command{
	Use:   "replace-string [string to inputpath] [string to outputpath]",
	Short: `<inputpath> にあるファイルを受け取り、<outputpath> に <inputpath> の内容を逆にした新しいファイルを作成します。`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(2),cobra.MaximumNArgs(3)),
    Run: func(cmd *cobra.Command, args []string) {
        inputfile := args[0]
        outputfile := "output.txt"
        niidle := args[1]
        newstring := ""
        if len(args) == 3 {
            newstring = args[2]
        }

        data, err := os.ReadFile(inputfile)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

        // データを置換する
        contents := string(data)
        contents = strings.ReplaceAll(contents,niidle,newstring)


        err = os.WriteFile(outputfile, []byte(contents), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
			os.Exit(1)
		}
    },
}

func init() {
	rootCmd.AddCommand(replace_stringCmd)
}
