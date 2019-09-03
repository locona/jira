package issue

import (
	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func NewCommandUpdate() *cobra.Command {
	// updateOption := &issue.Search{}
	cmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Update()
		},
	}

	return cmd
}

type UpdateCommand struct {
	Value  *jira.Issue
	Result []jira.Issue
}

func (cmd *UpdateCommand) Request(s *spinner.Spinner) error {
	updated, err := issue.Update(cmd.Value)
	if err != nil {
		return err
	}
	cmd.Result = []jira.Issue{*updated}
	return nil
}

func (cmd *UpdateCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

func Update() error {
	selectedIssue, err := selectIssue()
	if err != nil {
		return err
	}
	selectField := selectField()

	v := &jira.Issue{
		Key:    selectedIssue.Key,
		Fields: &jira.IssueFields{},
	}
	switch selectField {
	case issue.FieldSummary:
		summary := inputValue(selectedIssue.Fields.Summary)
		v.Fields.Summary = summary
	case issue.FieldDescription:
		description := inputValue(selectedIssue.Fields.Description)
		v.Fields.Description = description
	}

	return prompt.Progress(&UpdateCommand{
		Value: v,
	})
}

func selectIssue() (*jira.Issue, error) {
	issueList, err := issue.List(&issue.Search{
		Assignee: "own",
	})
	if err != nil {
		return nil, err
	}

	options, mapOptionToIssue := issue.Options(issueList)
	p := &survey.Select{
		Message: "Select Update ISSUE ID",
		Options: options,
	}
	var res string
	err = survey.AskOne(p, &res, nil)
	if err != nil {
		return nil, err
	}
	selected := mapOptionToIssue[res]
	return &selected, nil
}

func selectField() string {
	p := &survey.Select{
		Message: "Select Update Field",
		Options: issue.Fields,
	}
	var res string
	survey.AskOne(p, &res, nil)
	return res
}

func inputValue(current string) string {
	p := &survey.Multiline{
		Message: "Input Value",
		Default: current,
	}
	var res string
	survey.AskOne(p, &res, nil)
	return res
}
