package user

import (
	"fmt"
	"strings"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
)

type SearchUserList []SearchUser

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

	endpoint := "/rest/api/3/users/search"
	req, err := cli.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	list := SearchUserList{}
	_, err = cli.Do(req, &list)
	if err != nil {
		return nil, err
	}

	accountIDSlice := make([]string, 0)
	for _, user := range list {
		accountIDSlice = append(accountIDSlice, user.AccountID)
	}
	res, err := Bulk(accountIDSlice)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Bulk(accountIDSlice []string) ([]jira.User, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	res := make([]jira.User, 0)
	isLast := false

	queryAccountIDSlice := make([]string, 0)
	for _, accountID := range accountIDSlice {
		queryAccountIDSlice = append(queryAccountIDSlice, fmt.Sprintf("accountId=%v", accountID))
	}
	q := strings.Join(queryAccountIDSlice, "&")
	endpoint := fmt.Sprintf("/rest/api/3/user/bulk?%v", q)
	for !isLast {
		list := &UserList{}
		req, err := cli.NewRequest("GET", endpoint, nil)
		if err != nil {
			return nil, err
		}

		_, err = cli.Do(req, &list)
		if err != nil {
			return nil, err
		}

		res = append(res, list.Users...)

		isLast = list.IsLast
		endpoint = list.NextPage
	}

	return res, nil
}
