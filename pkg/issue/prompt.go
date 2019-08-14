package issue

import (
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
)

func ViewTable(issues []jira.Issue) {
	header := []string{"Key", "Summary", "Description"}
	data := make([][]string, len(issues))
	for i, _ := range issues {
		issue := issues[i]
		data[i] = []string{issue.Key, issue.Fields.Summary, issue.Fields.Description}
	}

	prompt.Table(header, data)
}
