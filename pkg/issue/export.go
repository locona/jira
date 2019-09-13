package issue

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/andygrunwald/go-jira"
	"gopkg.in/yaml.v2"
)

func Export(issueSlice []*jira.Issue) error {
	values := make([]*ApplyValue, len(issueSlice))
	for idx := range issueSlice {
		issue := issueSlice[idx]
		values[idx] = &ApplyValue{
			Key:         issue.Key,
			Summary:     issue.Fields.Summary,
			Description: issue.Fields.Description,
			Labels:      issue.Fields.Labels,
			Assignee:    issue.Fields.Assignee.Name,
			Type:        issue.Fields.Type.Name,
		}
	}

	out, err := yaml.Marshal(&values)
	if err != nil {
		return err
	}

	fname := fmt.Sprintf("jira-export-%v.yaml", time.Now().UnixNano())
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		return err
	}
	return nil
}
