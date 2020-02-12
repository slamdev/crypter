package internal

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:              "version",
	Short:            "Print the version number of " + rootCmdName,
	Long:             "All software has versions. This is " + rootCmdName + "'s",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("%v v0.1 -- HEAD\n", rootCmdName)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
