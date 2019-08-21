package user

import (
	"fmt"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
	"github.com/andygrunwald/go-jira"
)

type SearchUserList []*jira.User

type SearchUser struct {
	Self        string `json:"self"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	AccountType string `json:"accountType"`
}

type UserList struct {
	MaxResults int         `json:"maxResults"`
	StartAt    int         `json:"startAt"`
	Total      int         `json:"total"`
	IsLast     bool        `json:"isLast"`
	NextPage   string      `json:"nextPage"`
	Users      []jira.User `json:"values"`
}

func List() ([]jira.User, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	current, _ := project.Current()

	endpoint := fmt.Sprintf("/rest/api/3/user/assignable/multiProjectSearch?projectKeys=%v", current)
	req, err := cli.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	list := make([]jira.User, 0)
	_, err = cli.Do(req, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
