package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

func dbFilePath() string {
	return filepath.Join(xdg.StateHome, "clipboy", "clipboy.db")
}

var dbFileCmd = &cobra.Command{
	Use:   "db-file",
	Short: "Print the clipboard database path",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), dbFilePath())
		return err
	},
}

func init() {
	rootCmd.AddCommand(dbFileCmd)
}
