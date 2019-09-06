package issue

import (
	"log"

	"github.com/locona/jira/pkg/auth"
)

func BatchDelete(issueIDList []string) {
	for _, issueID := range issueIDList {
		err := Delete(issueID)
		if err != nil {
			log.Println(err)
		}
	}
}

func Delete(issueID string) error {
	cli, err := auth.Client()
	if err != nil {
		return err
	}

	_, err = cli.Issue.Delete(issueID)
	if err != nil {
		return err
	}
	return nil
}
