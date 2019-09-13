package issue

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func Label(i jira.Issue) string {
	assignee := "Undefined"
	if i.Fields.Assignee != nil {
		assignee = i.Fields.Assignee.Name
	}
	format := "%v(%v:%v): %v"
	return fmt.Sprintf(format, i.Key, assignee, i.Fields.Status.Name, i.Fields.Summary)
}
