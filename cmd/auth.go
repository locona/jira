package cmd

import (
	"github.com/locona/jira/cmd/project"
	"github.com/locona/jira/pkg/auth"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use: "auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Auth()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func Auth() error {
	promptBaseURL := promptui.Prompt{
		Label: "BaseURL",
	}
	baseURL, err := promptBaseURL.Run()
	if err != nil {
		return err
	}

	promptUsername := promptui.Prompt{
		Label: "User Name",
	}
	username, err := promptUsername.Run()
	if err != nil {
		return err
	}

	promptPassword := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, err := promptPassword.Run()
	if err != nil {
		return err
	}

	_auth, err := auth.Authenticate(baseURL, username, password)
	if err != nil {
		return err
	}

	err = _auth.Store()
	if err != nil {
		return err
	}

	err = project.Namespace()
	if err != nil {
		return err
	}

	return nil
}
