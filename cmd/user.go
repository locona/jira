package cmd

import (
	"github.com/3-shake/jira/cmd/user"
	"github.com/spf13/cobra"
)

// userCmd represents the output command
var userCmd = &cobra.Command{
	Use: "user",
}

func init() {
	listCmd := user.NewCommandList()
	showCmd := user.NewCommandShow()
	userCmd.AddCommand(
		listCmd,
		showCmd,
	)
	rootCmd.AddCommand(userCmd)
}
