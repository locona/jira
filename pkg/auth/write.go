package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Write(username, password string) error {
	auth := New(username, password)
	b, err := json.MarshalIndent(auth, "", " ")
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
