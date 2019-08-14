package issue

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func Label(i jira.Issue) string {
	format := "%v: %v"
	return fmt.Sprintf(format, i.Key, i.Fields.Summary)
}
