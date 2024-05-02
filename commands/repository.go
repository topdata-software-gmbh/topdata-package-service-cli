package commands

import (
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Commands related to repositories",
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
}
