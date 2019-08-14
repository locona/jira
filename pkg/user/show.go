package user

import (
	"github.com/3-shake/jira/pkg/auth"
	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

func FirstByEmail(email string) (*jira.User, error) {
	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	users, _, err := cli.User.Find(email)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.Wrapf(err, "Email: %v", email)
	}

	return &users[0], nil
}
