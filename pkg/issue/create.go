package issue

import (
	"log"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
	"github.com/3-shake/jira/pkg/user"
	"github.com/andygrunwald/go-jira"
)

type CreateValue struct {
	Summary     string `yaml:"summary,omitempty"`
	Description string `yaml:"description,omitempty"`
	// Epic        string   `yaml:"epic,omitempty"`
	Labels   []string       `yaml:"labels,omitempty"`
	Assignee string         `yaml:"assignee,omitempty"`
	Type     string         `yaml:"issuetype,omitempty"`
	Subtasks []*CreateValue `yaml:"subtasks,omitempty"`
}

func Create(v *CreateValue) ([]jira.Issue, error) {
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

	createdParent, _, err := cli.Issue.Create(&jira.Issue{
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
			// Epic: &jira.Epic{
			// Name: v.Epic,
			// },
		},
	})
	if err != nil {
		return nil, err
	}

	parentIssue, err := Show(createdParent.Key)
	if err != nil {
		return nil, err
	}

	subtasks := make([]*jira.Issue, 0)
	for idx := range v.Subtasks {
		subtasks = append(subtasks, &jira.Issue{
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
					ID: parentIssue.ID,
				},
				// Epic: &jira.Epic{
				// Name: v.Epic,
				// },
			},
		})
	}

	res := make([]jira.Issue, 0)
	res = append(res, *parentIssue)
	for idx := range subtasks {
		createdChild, _, err := cli.Issue.Create(subtasks[idx])
		if err != nil {
			log.Println(err)
		}

		childIssue, err := Show(createdChild.Key)
		if err != nil {
			return nil, err
		}
		res = append(res, *childIssue)
	}

	return res, nil
}
