package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching all repositories ...")
		response, err := topdataPackageServiceClient.ListRepositories()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", response)
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
}
