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
	applyCmd := issue.NewCommandApply()
	epicCmd := issue.NewCommandEpic()
	deleteCmd := issue.NewCommandDelete()
	inputCmd.AddCommand(
		listCmd,
		applyCmd,
		epicCmd,
		deleteCmd,
	)
	rootCmd.AddCommand(inputCmd)
}
