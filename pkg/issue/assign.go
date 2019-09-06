package issue

import (
	"github.com/locona/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
)

func UpdateAssignee(issueID string, assignee *jira.User) (*jira.Issue, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	_, err = cli.Issue.UpdateAssignee(issueID, &jira.User{AccountID: assignee.AccountID})
	if err != nil {
		return nil, err
	}

	res, err := Show(issueID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
