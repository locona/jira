package issue

import (
	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
	"github.com/3-shake/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
)

type CreateValue struct {
	Summary     string `yaml:"summary,omitempty"`
	Description string `yaml:"description,omitempty"`
	// Epic        string   `yaml:"epic,omitempty"`
	Labels   []string `yaml:"labels,omitempty"`
	Assignee string   `yaml:"assignee,omitempty"`
	Type     string   `yaml:"issuetype,omitempty"`
	// Subtasks    []*jira.Subtasks `yaml:"subtasks,omitempty"`

	project  string
	reporter string

	fields *jira.IssueFields
}

func Create(v *CreateValue) (*jira.Issue, error) {
	current, err := project.Current()
	if err != nil {
		return nil, err
	}

	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	myInfo, _ := auth.Read()
	reporter, err := user.FirstByEmail(myInfo.Username)
	if err != nil {
		return nil, err
	}

	assignee, err := user.FirstByEmail(v.Assignee)
	if err != nil {
		return nil, err
	}

	created, _, err := cli.Issue.Create(&jira.Issue{
		Fields: &jira.IssueFields{
			Summary:     v.Summary,
			Description: v.Description,
			Assignee:    assignee,
			Reporter:    reporter,
			Type: jira.IssueType{
				ID: "10033",
			},
			Project: jira.Project{
				Key: current,
			},
			// Epic: &jira.Epic{
			// Name: v.Epic,
			// },
			Labels: v.Labels,
		},
	})

	return Show(created.Key)
}
