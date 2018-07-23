package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kime",
	Long:  "Kime is an inventory inspector for OpenShift",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kime Inventory Inspector for OpenShift v0.1")
	},
}
