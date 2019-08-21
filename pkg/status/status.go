package status

import (
	"fmt"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/3-shake/jira/pkg/project"
)

type StatusList []struct {
	Self     string   `json:"self"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Subtask  bool     `json:"subtask"`
	Statuses []Status `json:"statuses"`
}

type Status struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
}

func List() (*StatusList, error) {
	projectName, err := project.Current()
	if err != nil {
		return nil, err
	}

	cli, err := auth.Client()
	if err != nil {
		return nil, err
	}

	statusList := new(StatusList)

	endpointFormat := "rest/api/2/project/%v/statuses"
	endpoint := fmt.Sprintf(endpointFormat, projectName)
	req, _ := cli.NewRequest("GET", endpoint, nil)
	_, err = cli.Do(req, statusList)
	if err != nil {
		return nil, err
	}

	return statusList, nil
}
