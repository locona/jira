package user

import (
	"github.com/locona/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
)

func ViewTable(users []jira.User) {
	header := []string{"Name", "Email Address"}
	data := make([][]string, len(users))
	for i, _ := range users {
		user := users[i]
		data[i] = []string{user.Name, user.EmailAddress}
	}

	prompt.Table(header, data)
}
