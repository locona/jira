package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/locona/jira/pkg/prompt"
)

func ViewTable(issues []jira.Issue) {
	header := []string{"ID", "Summary"}
	data := make([][]string, len(issues))
	for i, _ := range issues {
		issue := issues[i]
		data[i] = []string{issue.Key, issue.Fields.Summary}
	}

	prompt.Table(header, data)
}
