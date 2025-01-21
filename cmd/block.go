package cmd

import (
	"github.com/rende11/blocker/pkg"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "The block subcommand will add URL to blocked list.",
	Run: func(cmd *cobra.Command, args []string) {
		hostFilePath, _ := cmd.Flags().GetString("hosts-file")
		sites := GetSites(cmd, args)
		pkg.BlockSites(sites, hostFilePath)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
