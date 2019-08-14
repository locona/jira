package project

import (
	"fmt"
	"strings"

	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
)

func ViewTable(project *jira.Project) {
	header := []string{"Key", "Name", "Description", "IssueTypes"}
	issueTypeIDNameSlice := make([]string, 0)
	for _, issueType := range project.IssueTypes {
		issueTypeIDNameSlice = append(issueTypeIDNameSlice, fmt.Sprintf("%v: %v", issueType.ID, issueType.Name))
	}
	issueTypeIDName := strings.Join(issueTypeIDNameSlice, "\n")

	data := [][]string{
		[]string{project.Key, project.Name, project.Description, issueTypeIDName},
	}
	prompt.Table(header, data)
}
