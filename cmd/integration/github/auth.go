package github

import (
	"github.com/3-shake/jira/pkg/integration/github"
	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewCommandAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Auth()
		},
	}

	return cmd
}

type AuthCommand struct {
	Auth *github.Auth
}

func (cmd *AuthCommand) Request(s *spinner.Spinner) error {
	return nil
}

func (cmd *AuthCommand) Response() error {
	return nil
}

func Auth() error {
	tokenPrompt := promptui.Prompt{
		Label: "Token",
		Mask:  '*',
	}
	token, err := tokenPrompt.Run()
	if err != nil {
		return err
	}

	a, err := github.Authenticate(token)
	if err != nil {
		return err
	}

	err = a.Store()
	if err != nil {
		return err
	}
	return nil
}
