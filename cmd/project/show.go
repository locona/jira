package project

import (
	"github.com/3-shake/jira/pkg/project"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func NewCommandShow() *cobra.Command {
	cmd := &cobra.Command{
		Use: "show",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Show()
		},
	}

	return cmd
}

type ShowCommand struct {
	Result *jira.Project
}

func (cmd *ShowCommand) Request(s *spinner.Spinner) error {
	project, err := project.Show()
	if err != nil {
		return err
	}
	cmd.Result = project
	return nil
}

func (cmd *ShowCommand) Response() error {
	project.ViewTable(cmd.Result)
	return nil
}

func Show() error {
	return prompt.Loading(&ShowCommand{})
}
