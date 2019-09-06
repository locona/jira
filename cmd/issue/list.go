package issue

import (
	"github.com/locona/jira/pkg/issue"
	"github.com/locona/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func NewCommandList() *cobra.Command {
	listOption := &issue.Search{}
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(listOption)
		},
	}

	cmd.Flags().StringSliceVar(&listOption.Labels, "labels", []string{}, "Labels.")
	cmd.Flags().StringVar(&listOption.Status, "status", "", "Status.")
	cmd.Flags().StringVar(&listOption.Summary, "summary", "", "Summary.")
	cmd.Flags().StringVar(&listOption.Assignee, "assignee", "own", "Assignee.")
	cmd.Flags().StringVar(&listOption.Reporter, "reporter", "", "Reporter.")
	return cmd
}

type ListCommand struct {
	Option *issue.Search
	Result []jira.Issue
}

func (cmd *ListCommand) Request(s *spinner.Spinner) error {
	issues, err := issue.List(cmd.Option)
	if err != nil {
		return err
	}
	cmd.Result = issues
	return nil
}

func (cmd *ListCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

func List(option *issue.Search) error {
	return prompt.Progress(&ListCommand{
		Option: option,
	})
}
