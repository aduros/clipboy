package cmd

import (
    "os"
    "database/sql"

    "github.com/adrg/xdg"
    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

func openDb () *sql.DB {
    db, err := sql.Open("sqlite3", xdg.CacheHome+"/clipboy.db")
    if err != nil {
        panic(err)
    }
    return db
}

var listCmd = &cobra.Command {
    Use: "list",
    Short: "Print null-separated clipboard history",

    Run: func (cmd *cobra.Command, args []string) {
        db := openDb()
        defer db.Close()

        rows, err := db.Query("SELECT text FROM history ORDER BY rowid DESC")
        if err != nil {
            return // Don't print anything if the db doesn't exist yet
        }
        defer rows.Close()

        for rows.Next() {
            var text []byte
            err := rows.Scan(&text)
            if err != nil {
                panic(err)
            }
            os.Stdout.Write(text)
            os.Stdout.Write([]byte { 0 })
        }
    },
}

func init () {
    rootCmd.AddCommand(listCmd)
}
