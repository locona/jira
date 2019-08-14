package user

import (
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/3-shake/jira/pkg/user"
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
	user, err := user.FirstByEmail("miyamae@3-shake.com")
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
