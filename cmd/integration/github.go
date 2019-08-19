package integration

import (
	"github.com/3-shake/jira/cmd/integration/github"
	"github.com/spf13/cobra"
)

func NewCommandGithub() *cobra.Command {
	cmd := &cobra.Command{
		Use: "github",
	}

	authCmd := github.NewCommandAuth()
	pullrequestCmd := github.NewCommandPullRequest()
	cmd.AddCommand(authCmd, pullrequestCmd)
	return cmd
}
