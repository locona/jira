package project

import (
	"github.com/3-shake/jira/pkg/project"
	"github.com/spf13/cobra"
)

func NewCommandList() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return List()
		},
	}

	return cmd
}

func List() error {
	_, err := project.List()
	if err != nil {
		return err
	}

	return nil
}
