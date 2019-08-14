package issue

import (
	"fmt"
	"log"
	"time"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
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
	cmd.Flags().StringVar(&deleteOption.Search.Reporter, "reporter", "", "Status.")
	return cmd
}

func Delete(option *DeleteOption) error {
	if !option.Interactive {
		// BatchDelete(option.IssueIDList)
		return nil
	}

	listCmd := &ListCommand{Option: option.Search}
	err := prompt.Loading(listCmd)
	if err != nil {
		return err
	}

	issueList := listCmd.Result

	var selectedItem string
	prompt := &survey.Select{
		Message: "Select Delete Type",
		Options: []string{"SELECT", "CANCEL"},
	}

	err = survey.AskOne(prompt, &selectedItem)
	if err != nil {
		return err
	}

	if selectedItem == "CANCEL" {
		return nil
	}

	options := make([]string, 0)
	mapOptionToIssue := make(map[string]jira.Issue)
	for _, is := range issueList {
		op := issue.Label(is)
		options = append(options, op)
		mapOptionToIssue[op] = is
	}

	if selectedItem == "SELECT" {
		deletePrompt := &survey.MultiSelect{
			Message: "Select Delete ISSUE ID",
			Options: options,
		}
		deleteIssueOptions := make([]string, 0)
		err := survey.AskOne(deletePrompt, &deleteIssueOptions)
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
		survey.AskOne(prompt, &confirm)
		if !confirm {
			return nil
		}

		BatchDelete(deleteIssueSlice)
		return nil
	}

	// BatchDelete(option.IssueIDList)
	return nil
}

type DeleteCommand struct {
	Issue jira.Issue
}

func (cmd *DeleteCommand) Request(s *spinner.Spinner) error {
	label := issue.Label(cmd.Issue)
	s.Suffix = fmt.Sprintf("  %v", label)
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
	s := spinner.New(spinner.CharSets[3], 100*time.Millisecond) // Build our new spinner
	s.Color("magenta")
	for _, issue := range issueSlice {
		err := prompt.Loading(&DeleteCommand{
			Issue: issue,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
