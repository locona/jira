package cmd

import (
	"github.com/locona/jira/pkg/status"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use: "status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Status()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func Status() error {
	list, err := status.List()
	if err != nil {
		return err
	}

	pp.Println(list)
	return nil
}
