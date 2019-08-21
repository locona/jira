package issue

import (
	"log"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
	"github.com/3-shake/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
)

type ApplyValue struct {
	Key         string `yaml:"key,omitempty"`
	Summary     string `yaml:"summary,omitempty"`
	Description string `yaml:"description,omitempty"`
	// Epic        string   `yaml:"epic,omitempty"`
	Labels   []string      `yaml:"labels,omitempty"`
	Assignee string        `yaml:"assignee,omitempty"`
	Type     string        `yaml:"issuetype,omitempty"`
	Subtasks []*ApplyValue `yaml:"subtasks,omitempty"`
}

func Apply(v *ApplyValue) ([]jira.Issue, error) {
	current, err := project.Current()
	if err != nil {
		return nil, err
	}

	myInfo, _ := auth.Read()

	reporter, err := user.FirstByUsername(myInfo.Username)
	if err != nil {
		return nil, err
	}

	assignee, err := user.FirstByUsername(v.Assignee)
	if err != nil {
		return nil, err
	}

	parentIssue := &jira.Issue{
		Key: v.Key,
		Fields: &jira.IssueFields{
			Summary:     v.Summary,
			Description: v.Description,
			Labels:      v.Labels,
			Assignee:    assignee,
			Reporter:    reporter,
			Type: jira.IssueType{
				ID: "10033",
			},
			Project: jira.Project{
				Key: current,
			},
		},
	}

	appliedParentIssue, err := apply(parentIssue)
	if err != nil {
		return nil, err
	}

	subtasks := make([]*jira.Issue, 0)
	for idx := range v.Subtasks {
		subtasks = append(subtasks, &jira.Issue{
			Key: v.Subtasks[idx].Key,
			Fields: &jira.IssueFields{
				Summary:     v.Subtasks[idx].Summary,
				Description: v.Subtasks[idx].Description,
				Labels:      v.Subtasks[idx].Labels,
				Assignee:    assignee,
				Reporter:    reporter,
				Type: jira.IssueType{
					ID: "10058",
				},
				Project: jira.Project{
					Key: current,
				},
				Parent: &jira.Parent{
					ID: appliedParentIssue.ID,
				},
			},
		})
	}

	res := make([]jira.Issue, 0)
	res = append(res, *appliedParentIssue)
	for idx := range subtasks {
		appliedChildIssue, err := apply(subtasks[idx])
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, *appliedChildIssue)
	}

	return res, nil
}

func apply(issue *jira.Issue) (*jira.Issue, error) {
	if issue.Key == "" {
		return Create(issue)
	}

	return Update(issue)
}

func Create(issue *jira.Issue) (*jira.Issue, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	created, _, err := cli.Issue.Create(issue)
	if err != nil {
		return nil, err
	}

	res, err := Show(created.Key)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Update(issue *jira.Issue) (*jira.Issue, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	updated, _, err := cli.Issue.Update(issue)
	if err != nil {
		return nil, err
	}

	res, err := Show(updated.Key)
	if err != nil {
		return nil, err
	}

	return res, nil
}
