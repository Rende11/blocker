package cmd

import (
	"fmt"

	"github.com/rende11/blocker/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows hosts file content",
	Run: func(cmd *cobra.Command, args []string) {
		hostFilePath, _ := cmd.Flags().GetString("hosts-file")
		fmt.Println(pkg.GetHosts(hostFilePath))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
