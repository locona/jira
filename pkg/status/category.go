package status

import (
	"github.com/3-shake/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
)

func CategoryList() ([]jira.StatusCategory, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	list, _, err := cli.StatusCategory.GetList()
	if err != nil {
		return nil, err
	}

	return list, nil
}
