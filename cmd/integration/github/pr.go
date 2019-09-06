package github

import (
	cmdIssue "github.com/locona/jira/cmd/issue"
	"github.com/locona/jira/pkg/integration/github"
	"github.com/locona/jira/pkg/issue"
	"github.com/locona/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	ggithub "github.com/google/go-github/v27/github"
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

type PullRequestCommand struct {
	Value  *issue.ApplyValue
	Result []jira.Issue
}

func (cmd *PullRequestCommand) Request(s *spinner.Spinner) error {
	results, err := issue.Apply(cmd.Value)
	if err != nil {
		return err
	}

	cmd.Result = results
	return nil
}

func (cmd *PullRequestCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

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
	if err != nil {
		return err
	}

	assignee, err := cmdIssue.SelectAssignee()
	if err != nil {
		return err
	}

	subtasks := make([]*issue.ApplyValue, 0)
	for idx := range commits {
		subtasks = append(subtasks, &issue.ApplyValue{
			Summary:  *commits[idx].Commit.Message,
			Assignee: assignee.Name,
			Labels:   labels,
		})
	}
	return prompt.Progress(&PullRequestCommand{
		Value: &issue.ApplyValue{
			Summary:  *selectedPR.Title,
			Labels:   labels,
			Assignee: assignee.Name,
			Subtasks: subtasks,
		},
	})

	return nil
}
