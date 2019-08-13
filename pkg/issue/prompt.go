package issue

import "github.com/3-shake/jira/pkg/prompt"

func ViewTable(list []*Issue) {
	header := []string{"Key", "Summary", "Description"}
	data := make([][]string, 0)
	for _, issue := range list {
		d := []string{issue.Key, issue.Fields.Summary, issue.Fields.Description}
		data = append(data, d)
	}

	prompt.Table(header, data)
}
