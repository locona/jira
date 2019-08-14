package user

import (
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/3-shake/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
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

type ListCommand struct {
	Result []jira.User
}

func (cmd *ListCommand) Request(s *spinner.Spinner) error {
	users, err := user.List()
	if err != nil {
		return err
	}

	cmd.Result = users
	return nil
}

func (cmd *ListCommand) Response() error {
	user.ViewTable(cmd.Result)
	return nil
}

func List() error {
	return prompt.Progress(&ListCommand{})
}
