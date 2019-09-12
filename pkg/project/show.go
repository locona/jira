package project

import (
	"github.com/andygrunwald/go-jira"
	"github.com/locona/jira/pkg/auth"
)

func Show() (*jira.Project, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	current, _ := Current()
	project, _, err := cli.Project.Get(current)
	if err != nil {
		return nil, err
	}

	return project, nil
}
