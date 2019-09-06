package project

import (
	"github.com/locona/jira/pkg/project"
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

//
// type ListCommand struct {
// Result *jira.Project
// }
//
// func (cmd *ListCommand) Request(s *spinner.Spinner) error {
// project, err := project.Show()
// if err != nil {
// return err
// }
// cmd.Result = project
// return nil
// }
//
// func (cmd *ListCommand) Response() error {
// project.ViewTable(cmd.Result)
// return nil
// }

func List() error {
	// return prompt.Loading(&ShowCommand{})
	_, err := project.List()
	if err != nil {
		return err
	}

	return nil
}
