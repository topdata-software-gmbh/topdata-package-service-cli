package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := topdataPackageServiceClient.Ping()
		if err != nil {
			fmt.Printf("Error pinging server: %v\n", err)
			return
		}
		fmt.Printf("Server response: %s\n", response)
	}}

func init() {
	rootCmd.AddCommand(pingCmd)
}
