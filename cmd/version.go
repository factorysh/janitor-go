package cmd

import (
	"fmt"

	"github.com/factorysh/janitor-go/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of janitor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version())
	},
}
