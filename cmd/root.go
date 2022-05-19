package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "clipboy",
    Short: "Clipboy is a minimal clipboard manager",
}

func Execute () {
    err := rootCmd.Execute()
    if err != nil {
        panic(err)
    }
}
