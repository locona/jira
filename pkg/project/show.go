package project

import (
	"github.com/locona/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
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
