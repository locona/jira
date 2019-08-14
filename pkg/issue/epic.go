package issue

import (
	"fmt"
	"strings"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
	"github.com/andygrunwald/go-jira"
)

func Epic() ([]jira.Issue, error) {
	projectName, err := project.Current()
	if err != nil {
		return nil, err
	}

	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	jqlList := make([]string, 0)
	jqlList = append(jqlList, fmt.Sprintf("project = %v", projectName))
	jqlList = append(jqlList, fmt.Sprintf("issuetype = %v", "Epic"))

	jql := strings.Join(jqlList, " AND ")
	issues, _, err := cli.Issue.Search(jql, nil)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
