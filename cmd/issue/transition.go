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

func NewCommandTransition() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transition",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Transition()
		},
	}

	return cmd
}

type TransitionCommand struct {
	Issue      *jira.Issue
	Transition *jira.Transition
}

func (cmd *TransitionCommand) Request(s *spinner.Spinner) error {
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

func (cmd *TransitionCommand) Response() error {
	return nil
}

func Transition() error {
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
		err := prompt.Progress(&TransitionCommand{
			Issue:      selectedIssueSlice[idx],
			Transition: selectedTransition,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

//
// func multiSelectIssue() ([]*jira.Issue, error) {
// issueList, err := issue.List(&issue.Search{})
// if err != nil {
// return nil, err
// }
//
// options := make([]string, 0)
// mapOptionToIssue := make(map[string]jira.Issue)
// for _, is := range issueList {
// op := issue.Label(is)
// options = append(options, op)
// mapOptionToIssue[op] = is
// }
//
// prompt := &survey.MultiSelect{
// Message: "Select the issue status you whose status you want to change.",
// Options: options,
// }
// targetOptionSlice := make([]string, 0)
// err = survey.AskOne(prompt, &targetOptionSlice, nil)
// if err != nil {
// return nil, err
// }
//
// res := make([]*jira.Issue, 0)
// for _, target := range targetOptionSlice {
// is := mapOptionToIssue[target]
// res = append(res, &is)
// }
//
// return res, nil
// }

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
