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
	assignCmd := issue.NewCommandAssign()
	applyCmd := issue.NewCommandApply()
	epicCmd := issue.NewCommandEpic()
	deleteCmd := issue.NewCommandDelete()
	statusCmd := issue.NewCommandStatus()
	inputCmd.AddCommand(
		listCmd,
		assignCmd,
		applyCmd,
		epicCmd,
		deleteCmd,
		statusCmd,
	)
	rootCmd.AddCommand(inputCmd)
}
