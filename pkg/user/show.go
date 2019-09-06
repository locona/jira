package user

import (
	"github.com/locona/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

func FirstByUsername(username string) (*jira.User, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	users, _, err := cli.User.Find(username)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.Wrapf(err, "Username: %v", username)
	}

	return &users[0], nil
}
