package auth

import (
	"fmt"
	"os/user"
)

type auth struct {
	BaseURL  string `json:"base_url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(username, password string) *auth {
	return &auth{
		BaseURL:  "https://3-shake.atlassian.net",
		Username: username,
		Password: password,
	}
}

func confPath() string {
	user, _ := user.Current()
	path := fmt.Sprintf("%v/.config/jira/config.json", user.HomeDir)
	return path
}
