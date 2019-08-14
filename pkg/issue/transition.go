package issue

import (
	"github.com/3-shake/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
)

func TransitionList(issueID string) ([]jira.Transition, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	transitions, _, err := cli.Issue.GetTransitions(issueID)
	if err != nil {
		return nil, err
	}

	return transitions, nil
}

func ChangeTransition(issueID, transitionID string) error {
	cli, err := auth.Client()
	if err != nil {
		return err
	}

	_, err = cli.Issue.DoTransition(issueID, transitionID)
	if err != nil {
		return err
	}

	return nil
}
