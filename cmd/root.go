package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blocky",
	Short: "Blocking local access to websites that distract you.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "List of sites")

	rootCmd.PersistentFlags().StringP("hosts-file", "f", "", "Custom hosts file path other than /etc/hosts")
	rootCmd.MarkFlagFilename("hosts-file")
}
