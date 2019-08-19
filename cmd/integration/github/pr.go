package github

import (
	"github.com/3-shake/jira/pkg/integration/github"
	ggithub "github.com/google/go-github/v27/github"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

func NewCommandPullRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pr",
		RunE: func(cmd *cobra.Command, args []string) error {
			return PullRequest()
		},
	}

	return cmd
}

// type PullRequestCommand struct {
// Option *github.Search
// Result []jira.Issue
// }
//
// func (cmd *PullRequestCommand) Request(s *spinner.Spinner) error {
// githubs, err := github.PullRequest(cmd.Option)
// if err != nil {
// return err
// }
// cmd.Result = githubs
// return nil
// }
//
// func (cmd *PullRequestCommand) Response() error {
// github.ViewTable(cmd.Result)
// return nil
// }

func PullRequest() error {
	pullrequests, err := github.PullRequests()

	options := make([]string, 0)
	mapOptionToPR := make(map[string]ggithub.PullRequest)
	for _, pr := range pullrequests {
		options = append(options, pr.Title)
		mapOptionToPR[pr.Title] = &pr
	}

	prPrompt := &survey.Select{
		Message: "Select PR",
		Options: options,
	}
	var prOption string
	err = survey.AskOne(prPrompt, &prOption, nil)
	if err != nil {
		return err
	}

	selectedPR := mapOptionToPR[prOption]
	pp.Println(selectedPR)

	return nil
	// return prompt.Progress(&PullRequestCommand{
	// Option: option,
	// })
}
