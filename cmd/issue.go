package cmd

import (
	"github.com/3-shake/jira/cmd/issue"
	"github.com/spf13/cobra"
)

// issueCmd represents the output command
var issueCmd = &cobra.Command{
	Use: "issue",
}

func init() {
	listCmd := issue.NewCommandList()
	assignCmd := issue.NewCommandAssign()
	applyCmd := issue.NewCommandApply()
	epicCmd := issue.NewCommandEpic()
	updateCmd := issue.NewCommandUpdate()
	deleteCmd := issue.NewCommandDelete()
	statusCmd := issue.NewCommandStatus()
	issueCmd.AddCommand(
		listCmd,
		assignCmd,
		applyCmd,
		epicCmd,
		updateCmd,
		deleteCmd,
		statusCmd,
	)
	rootCmd.AddCommand(issueCmd)
}
