package user

import (
	"github.com/locona/jira/pkg/prompt"
	"github.com/locona/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/k0kubun/pp"
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
	Result []jira.User
}

func (cmd *ShowCommand) Request(s *spinner.Spinner) error {
	user, err := user.FirstByUsername("miyamae@locona.com")
	if err != nil {
		return err
	}
	pp.Println(user)
	// cmd.Result = users
	return nil
}

func (cmd *ShowCommand) Response() error {
	// user.ViewTable(cmd.Result)
	return nil
}

func Show() error {
	return prompt.Progress(&ShowCommand{})
}
