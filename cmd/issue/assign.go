package issue

import (
	"errors"
	"fmt"
	"log"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/3-shake/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func NewCommandAssign() *cobra.Command {
	cmd := &cobra.Command{
		Use: "assign",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Assign()
		},
	}

	return cmd
}

type AssignCommand struct {
	Issue    *jira.Issue
	Assignee *jira.User
}

func (cmd *AssignCommand) Request(s *spinner.Spinner) error {
	label := issue.Label(*cmd.Issue)

	suffixFormat := "%v : Status To `%v`"
	suffixMsg := fmt.Sprintf(suffixFormat, label, cmd.Assignee.EmailAddress)

	var suf = make([]byte, 100)
	copy(suf, suffixMsg)
	s.Suffix = string(suf)
	s.FinalMSG = fmt.Sprintf("%v  %v \n", prompt.IconClear, suffixMsg)

	_, err := issue.UpdateAssignee(cmd.Issue.ID, cmd.Assignee)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *AssignCommand) Response() error {
	return nil
}

func Assign() error {
	selectedIssueSlice, err := multiSelectIssue("Select the issue status you whose status you want to change.")
	if err != nil {
		return err
	}

	if len(selectedIssueSlice) < 1 {
		return errors.New("Required Select Issue")
	}

	selectedAssignee, err := selectAssignee()
	if err != nil {
		return err
	}

	for idx := range selectedIssueSlice {
		err := prompt.Progress(&AssignCommand{
			Issue:    selectedIssueSlice[idx],
			Assignee: selectedAssignee,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func selectAssignee() (*jira.User, error) {
	users, err := user.List()
	if err != nil {
		return nil, err
	}

	options := make([]string, 0)
	mapNameToAssignee := make(map[string]jira.User, 0)
	for _, u := range users {
		mapNameToAssignee[u.Name] = u
		options = append(options, u.Name)
	}

	prompt := &survey.Select{
		Message: "Select the assignee you want to change.",
		Options: options,
	}
	var targetAssignee string
	err = survey.AskOne(prompt, &targetAssignee, nil)
	if err != nil {
		return nil, err
	}
	res := mapNameToAssignee[targetAssignee]
	return &res, nil
}
