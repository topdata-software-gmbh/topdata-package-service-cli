package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
    Use:   "ping",
    Short: "Ping the server",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Pong!")
    },
}

func init() {
    rootCmd.AddCommand(pingCmd)
}
