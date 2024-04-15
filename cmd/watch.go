package cmd

import (
    "context"

    "github.com/spf13/cobra"
    "golang.design/x/clipboard"
)

var watchCmd = &cobra.Command {
    Use: "watch",
    Short: "Watch clipboard history",

    Run: func (cmd *cobra.Command, args []string) {
        db := openDb()
        defer db.Close()

        stmt := `CREATE TABLE IF NOT EXISTS history (
            time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            text TEXT NOT NULL
        );`

        _, err := db.Exec(stmt)
        if err != nil {
            panic(err)
        }

        ch := clipboard.Watch(context.Background(), clipboard.FmtText)
        for data := range ch {
            stmt, err := db.Prepare("INSERT INTO history (text) VALUES(?)")
            if err != nil {
                panic(err)
            }

            _, err = stmt.Exec(data)
            if err != nil {
                panic(err)
            }

            // Clean up old items
            db.Exec("DELETE FROM history WHERE time < datetime('now', '-7 days')")
        }
    },
}

func init () {
    rootCmd.AddCommand(watchCmd)
}
