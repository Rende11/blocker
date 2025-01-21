package cmd

import (
	"github.com/rende11/blocker/pkg"

	"github.com/spf13/cobra"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		hostFilePath, _ := cmd.Flags().GetString("hosts-file")
		sitesFromFlag, _ := cmd.Flags().GetStringSlice("sites")

		sitesFromArgs := args

		var sites []string

		if len(sitesFromArgs) == 0 {
			sites = sitesFromFlag
		} else {
			sites = sitesFromArgs
		}
		pkg.UnBlockSites(sites, hostFilePath)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}
