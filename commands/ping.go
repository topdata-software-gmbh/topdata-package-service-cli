package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pinging server ...")
		response, err := topdataPackageServiceClient.Ping()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Response: %s\n", response)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
