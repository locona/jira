package issue

import (
	"errors"
	"fmt"
	"log"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func NewCommandStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use: "status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Status()
		},
	}

	return cmd
}

type StatusCommand struct {
	Issue      *jira.Issue
	Transition *jira.Transition
}

func (cmd *StatusCommand) Request(s *spinner.Spinner) error {
	label := issue.Label(*cmd.Issue)

	suffixFormat := "%v : Status To `%v`"
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

func (cmd *StatusCommand) Response() error {
	return nil
}

func Status() error {
	selectedIssueSlice, err := multiSelectIssue("Select the issue status you whose status you want to change.")
	if err != nil {
		return err
	}

	if len(selectedIssueSlice) < 1 {
		return errors.New("Required Select Issue")
	}

	selectedTransition, err := selectTransition(selectedIssueSlice[0].Key)
	if err != nil {
		return err
	}

	for idx := range selectedIssueSlice {
		err := prompt.Progress(&StatusCommand{
			Issue:      selectedIssueSlice[idx],
			Transition: selectedTransition,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func selectTransition(issueID string) (*jira.Transition, error) {
	transitions, err := issue.TransitionList(issueID)
	if err != nil {
		return nil, err
	}

	options := make([]string, 0)
	mapNameToTransition := make(map[string]jira.Transition, 0)
	for _, transition := range transitions {
		mapNameToTransition[transition.Name] = transition
		options = append(options, transition.Name)
	}

	prompt := &survey.Select{
		Message: "Select the status you want to change.",
		Options: options,
	}
	var targetTransition string
	err = survey.AskOne(prompt, &targetTransition, nil)
	if err != nil {
		return nil, err
	}
	res := mapNameToTransition[targetTransition]
	return &res, nil
}
