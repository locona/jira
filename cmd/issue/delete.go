package issue

import (
	"fmt"
	"log"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type DeleteOption struct {
	Interactive bool
	IssueIDList []string

	Search *issue.Search
}

func NewCommandDelete() *cobra.Command {
	deleteOption := &DeleteOption{
		Search: &issue.Search{},
	}
	cmd := &cobra.Command{
		Use: "delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Delete(deleteOption)
		},
	}

	cmd.Flags().BoolVarP(&deleteOption.Interactive, "interactive", "i", false, "Interactive.")
	cmd.Flags().StringSliceVar(&deleteOption.IssueIDList, "issueid", []string{}, "Issue ID.")

	// search
	cmd.Flags().StringSliceVar(&deleteOption.Search.Labels, "labels", []string{}, "Labels.")
	cmd.Flags().StringVar(&deleteOption.Search.Status, "status", "", "Status.")
	cmd.Flags().StringVar(&deleteOption.Search.Summary, "summary", "", "Summary.")
	cmd.Flags().StringVar(&deleteOption.Search.Assignee, "assignee", "own", "Assignee.")
	cmd.Flags().StringVar(&deleteOption.Search.Reporter, "reporter", "", "Reporter.")
	return cmd
}

func Delete(option *DeleteOption) error {
	if !option.Interactive {
		// BatchDelete(option.IssueIDList)
		return nil
	}

	issueList, err := issue.List(option.Search)
	if err != nil {
		return err
	}

	options, mapOptionToIssue := issue.Options(issueList)
	deletePrompt := &survey.MultiSelect{
		Message: "Select Delete ISSUE ID",
		Options: options,
	}
	deleteIssueOptions := make([]string, 0)
	err = survey.AskOne(deletePrompt, &deleteIssueOptions, nil)
	if err != nil {
		return err
	}

	deleteIssueSlice := make([]jira.Issue, 0)
	for _, op := range deleteIssueOptions {
		issue := mapOptionToIssue[op]
		deleteIssueSlice = append(deleteIssueSlice, issue)
	}

	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you really want to delete this?",
	}
	survey.AskOne(prompt, &confirm, nil)
	if !confirm {
		return nil
	}

	BatchDelete(deleteIssueSlice)
	return nil
}

type DeleteCommand struct {
	Issue jira.Issue
}

func (cmd *DeleteCommand) Request(s *spinner.Spinner) error {
	label := issue.Label(cmd.Issue)

	var suf = make([]byte, 100)
	copy(suf, label)
	s.Suffix = string(suf)

	s.FinalMSG = fmt.Sprintf("%v  %v \n", prompt.IconClear, label)
	err := issue.Delete(cmd.Issue.Key)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *DeleteCommand) Response() error {
	return nil
}

func BatchDelete(issueSlice []jira.Issue) {
	for _, issue := range issueSlice {
		err := prompt.Progress(&DeleteCommand{
			Issue: issue,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
