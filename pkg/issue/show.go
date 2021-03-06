package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/locona/jira/pkg/auth"
)

func Show(issueID string) (*jira.Issue, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	issue, _, err := cli.Issue.Get(issueID, nil)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
