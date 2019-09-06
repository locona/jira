package issue

import (
	"github.com/locona/jira/pkg/issue"
	"github.com/andygrunwald/go-jira"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func multiSelectIssue(msg string) ([]*jira.Issue, error) {
	issueList, err := issue.List(&issue.Search{})
	if err != nil {
		return nil, err
	}

	options := make([]string, 0)
	mapOptionToIssue := make(map[string]jira.Issue)
	for _, is := range issueList {
		op := issue.Label(is)
		options = append(options, op)
		mapOptionToIssue[op] = is
	}

	prompt := &survey.MultiSelect{
		Message: "Select the issue status you whose status you want to change.",
		Options: options,
	}
	targetOptionSlice := make([]string, 0)
	err = survey.AskOne(prompt, &targetOptionSlice, nil)
	if err != nil {
		return nil, err
	}

	res := make([]*jira.Issue, 0)
	for _, target := range targetOptionSlice {
		is := mapOptionToIssue[target]
		res = append(res, &is)
	}

	return res, nil
}
