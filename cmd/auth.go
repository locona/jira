package cmd

import (
	"github.com/3-shake/jira/cmd/project"
	"github.com/3-shake/jira/pkg/auth"
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

	_auth, err := auth.Authenticate(username, password)
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
