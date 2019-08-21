package project

import (
	"github.com/3-shake/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
)

func List() (*jira.ProjectList, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	list, _, err := cli.Project.GetList()
	if err != nil {
		return nil, err
	}

	return list, nil
}
