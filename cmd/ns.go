package cmd

import (
	"github.com/locona/jira/cmd/project"
	"github.com/spf13/cobra"
)

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use: "ns",
	RunE: func(cmd *cobra.Command, args []string) error {
		return project.Namespace()
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
}
