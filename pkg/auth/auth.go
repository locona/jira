package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

type Auth struct {
	BaseURL  string `json:"base_url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) Health() error {
	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(a.Username),
		Password: strings.TrimSpace(a.Password),
	}

	cli, err := jira.NewClient(tp.Client(), strings.TrimSpace(a.BaseURL))
	if err != nil {
		return err
	}

	users, _, err := cli.User.Find(a.Username)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.Wrapf(err, "Email: %v", a.Username)
	}

	return err
}

func (a *Auth) Store() error {
	b, err := json.MarshalIndent(a, "", " ")
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

func Client() (*jira.Client, error) {
	a, err := Read()
	if err != nil {
		return nil, err
	}

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(a.Username),
		Password: strings.TrimSpace(a.Password),
	}

	cli, err := jira.NewClient(tp.Client(), strings.TrimSpace(a.BaseURL))
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func Authenticate(baseURL, username, password string) (*Auth, error) {
	auth := &Auth{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
	}

	err := auth.Health()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func confPath() string {
	user, _ := user.Current()
	path := fmt.Sprintf("%v/.config/jira/config.json", user.HomeDir)
	return path
}
