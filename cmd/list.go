package cmd

import (
    "database/sql"
    "os"
    "path/filepath"

    "github.com/adrg/xdg"
    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

func openDb () *sql.DB {
    dir := filepath.Join(xdg.StateHome, "clipboy")
    if err := os.MkdirAll(dir, 0700); err != nil {
        panic(err)
    }
    db, err := sql.Open("sqlite3", filepath.Join(dir, "clipboy.db"))
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
