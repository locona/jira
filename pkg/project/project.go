package project

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func Current() (string, error) {
	fp, err := os.Open(confPath())
	if err != nil {
		return "", err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	return scanner.Text(), nil
}

func Store(projectName string) error {
	confPath := confPath()
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		dir := filepath.Dir(confPath)
		os.MkdirAll(dir, 0755)
	}
	err := ioutil.WriteFile(confPath, []byte(projectName), 0644)
	if err != nil {
		return err
	}
	return nil
}

func confPath() string {
	user, _ := user.Current()
	path := fmt.Sprintf("%v/.config/jira/project", user.HomeDir)
	return path
}
