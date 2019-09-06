package issue

import (
	"github.com/locona/jira/pkg/issue"
	"github.com/locona/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func NewCommandEpic() *cobra.Command {
	cmd := &cobra.Command{
		Use: "epic",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Epic()
		},
	}
	return cmd
}

type EpicCommand struct {
	Result []jira.Issue
}

func (cmd *EpicCommand) BeforeRequest(s *spinner.Spinner) *spinner.Spinner {
	return s
}

func (cmd *EpicCommand) Request(s *spinner.Spinner) error {
	epics, err := issue.Epic()
	if err != nil {
		return err
	}
	cmd.Result = epics
	return nil
}

func (cmd *EpicCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

func Epic() error {
	return prompt.Progress(&EpicCommand{})
}
