package issue

import (
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
)

func ViewTable(issues []jira.Issue) {
	header := []string{"ID", "ParentID", "Summary", "Description"}
	data := make([][]string, len(issues))
	for i, _ := range issues {
		issue := issues[i]
		parentID := ""
		if issue.Fields.Parent != nil {
			parentID = issue.Fields.Parent.Key
		}
		data[i] = []string{issue.Key, parentID, issue.Fields.Summary, issue.Fields.Description}
	}

	prompt.Table(header, data)
}
