package cmd

import (
	"github.com/3-shake/jira/cmd/project"
	"github.com/spf13/cobra"
)

// inputCmd represents the output command
var projectCmd = &cobra.Command{
	Use: "project",
}

func init() {
	listCmd := project.NewCommandList()
	nsCmd := project.NewCommandNamespace()
	projectCmd.AddCommand(
		listCmd,
		nsCmd,
	)
	rootCmd.AddCommand(projectCmd)
}
