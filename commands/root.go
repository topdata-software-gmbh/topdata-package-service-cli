package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/pkg"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "topdata-package-service-cli",
	Short: "A CLI for the Topdata Package Service",
	Long: `topdata-package-service-cli is a CLI tool to interact with the Topdata Package Service

... some more text here ...
`,
}

var topdataPackageServiceClient *pkg.TopdataPackageServiceClient

func Execute(client *pkg.TopdataPackageServiceClient) {
	topdataPackageServiceClient = client
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
