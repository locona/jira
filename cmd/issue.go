package cmd

import (
	"github.com/3-shake/jira/cmd/issue"
	"github.com/spf13/cobra"
)

// inputCmd represents the output command
var inputCmd = &cobra.Command{
	Use: "issue",
}

func init() {
	listCmd := issue.NewCommandList()
	createCmd := issue.NewCommandCreate()
	epicCmd := issue.NewCommandEpic()
	deleteCmd := issue.NewCommandDelete()
	transitionCmd := issue.NewCommandTransition()
	inputCmd.AddCommand(
		listCmd,
		createCmd,
		epicCmd,
		deleteCmd,
		transitionCmd,
	)
	rootCmd.AddCommand(inputCmd)
}
