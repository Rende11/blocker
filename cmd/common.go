package cmd

import "github.com/spf13/cobra"

func GetSites(cmd *cobra.Command, args []string) []string {
	sitesFromFlag, _ := cmd.Flags().GetStringSlice("sites")
	sitesFromArgs := args

	var sites []string
	if len(sitesFromArgs) == 0 {
		sites = sitesFromFlag
	} else {
		sites = sitesFromArgs
	}
	return sites
}
