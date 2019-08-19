package cmd

import (
	"github.com/3-shake/jira/cmd/integration"
	"github.com/spf13/cobra"
)

// integrationCmd represents the output command
var integrationCmd = &cobra.Command{
	Use: "integration",
}

func init() {
	githubCmd := integration.NewCommandGithub()
	integrationCmd.AddCommand(
		githubCmd,
	)
	rootCmd.AddCommand(integrationCmd)
}
