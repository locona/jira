package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v27/github"
)

type Auth struct {
	Token string `json:"token"`
}

func (a *Auth) Health() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: a.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	_, _, err := client.Repositories.List(ctx, "", nil)
	return err
}

func (a *Auth) Store() error {
	b, err := json.MarshalIndent(a, "", "")
	if err != nil {
		return err
	}

	confPath := confPath()
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		dir := filepath.Dir(confPath)
		os.MkdirAll(dir, 0755)
	}
	err = ioutil.WriteFile(confPath, b, 0644)
	if err != nil {
		return err
	}

	return nil

}

func Read() (*Auth, error) {
	b, err := ioutil.ReadFile(confPath())
	if err != nil {
		return nil, err
	}

	a := &Auth{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func Client() *github.Client {
	a, _ := Read()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: a.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func Authenticate(token string) (*Auth, error) {
	auth := &Auth{
		Token: token,
	}

	err := auth.Health()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func confPath() string {
	user, _ := user.Current()
	path := fmt.Sprintf("%v/.config/jira/integration/github/config.json", user.HomeDir)
	return path
}
