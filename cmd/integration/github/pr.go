package github

import (
	"github.com/3-shake/jira/pkg/integration/github"
	"github.com/3-shake/jira/pkg/issue"
	ggithub "github.com/google/go-github/v27/github"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var labels = []string{"github-pr", "qa"}

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
		options = append(options, *pr.Title)
		mapOptionToPR[*pr.Title] = *pr
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
	commits, err := github.PullRequestCommits(*selectedPR.Number)

	subtasks := make([]*issue.ApplyValue, 0)
	for idx := range commits {
		subtasks = append(subtasks, &issue.ApplyValue{
			Summary: *commits[idx].Commit.Message,
			Labels:  labels,
		})
	}

	results, err := issue.Apply(&issue.ApplyValue{
		Summary:  *selectedPR.Title,
		Labels:   labels,
		Subtasks: subtasks,
	})

	pp.Println(results, err)

	return nil
	// return prompt.Progress(&PullRequestCommand{
	// Option: option,
	// })
}
