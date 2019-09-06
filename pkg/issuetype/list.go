package issuetype

import "github.com/locona/jira/pkg/auth"

func List() {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	issue, _, err := cli.Project.Get(issueID, nil)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
