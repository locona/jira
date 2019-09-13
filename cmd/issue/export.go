package issue

import (
	"errors"
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/locona/jira/pkg/issue"
	"github.com/locona/jira/pkg/prompt"
	"github.com/spf13/cobra"
)

func NewCommandExport() *cobra.Command {
	cmd := &cobra.Command{
		Use: "export",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Export()
		},
	}

	return cmd
}

type ExportCommand struct {
	Issue      *jira.Issue
	Transition *jira.Transition
}

func (cmd *ExportCommand) Request(s *spinner.Spinner) error {
	label := issue.Label(*cmd.Issue)

	suffixFormat := "%v : Export To `%v`"
	suffixMsg := fmt.Sprintf(suffixFormat, label, cmd.Transition.Name)

	var suf = make([]byte, 100)
	copy(suf, suffixMsg)
	s.Suffix = string(suf)

	s.FinalMSG = fmt.Sprintf("%v  %v \n", prompt.IconClear, suffixMsg)

	err := issue.ChangeTransition(cmd.Issue.Key, cmd.Transition.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *ExportCommand) Response() error {
	return nil
}

func Export() error {
	selectedIssueSlice, err := multiSelectIssue("Select the issue status you whose status you want to change.")
	if err != nil {
		return err
	}

	if len(selectedIssueSlice) < 1 {
		return errors.New("Required Select Issue")
	}

	err = issue.Export(selectedIssueSlice)
	if err != nil {
		return err
	}

	return nil
}
