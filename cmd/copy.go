package cmd

import (
    "os"
    "io/ioutil"

    "github.com/spf13/cobra"
    "golang.design/x/clipboard"
)

var copyCmd = &cobra.Command {
    Use: "copy",
    Short: "Copy stdin to the clipboard",

    Run: func (cmd *cobra.Command, args []string) {
        data, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            panic(err)
        }
        if len(data) > 0 {
            clipboard.Write(clipboard.FmtText, data)
        }
    },
}

func init () {
    rootCmd.AddCommand(copyCmd)
}
