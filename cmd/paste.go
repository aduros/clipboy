package cmd

import (
    "os"

    "github.com/spf13/cobra"
    "golang.design/x/clipboard"
)

var pasteCmd = &cobra.Command {
    Use: "paste",
    Short: "Paste the clipboard to stdout",

    Run: func (cmd *cobra.Command, args []string) {
        os.Stdout.Write(clipboard.Read(clipboard.FmtText))
    },
}

func init () {
    rootCmd.AddCommand(pasteCmd)
}
