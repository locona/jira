package issue

import (
	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
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

	cmd.Flags().BoolVarP(&listOption.Verbose, "verbose", "v", false, "Verbose.")
	cmd.Flags().StringSliceVar(&listOption.Labels, "labels", []string{}, "Labels.")
	cmd.Flags().StringVar(&listOption.Status, "status", "", "Status.")
	cmd.Flags().StringVar(&listOption.Summary, "summary", "", "Summary.")
	cmd.Flags().StringVar(&listOption.Assignee, "assignee", "own", "Assignee.")
	cmd.Flags().StringVar(&listOption.Reporter, "reporter", "", "Reporter.")
	return cmd
}

type ListCommand struct {
	Option *issue.Search
	Result []*issue.Issue
}

func (cmd *ListCommand) Request(s *spinner.Spinner) error {
	list, err := issue.List(cmd.Option)
	if err != nil {
		return err
	}
	cmd.Result = list
	return nil
}

func (cmd *ListCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

func List(option *issue.Search) error {
	return prompt.Loading(&ListCommand{
		Option: option,
	})
}
