package issue

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type CreateOption struct {
	FilePath string
}

func NewCommandCreate() *cobra.Command {
	createOption := &CreateOption{}
	cmd := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Create(createOption)
		},
	}

	cmd.Flags().StringVarP(&createOption.FilePath, "file", "f", "", "Status.")
	return cmd
}

type CreateCommand struct {
	Value *issue.CreateValue
}

func (cmd *CreateCommand) Request(s *spinner.Spinner) error {
	// NOTE: for bug survey
	var suf = make([]byte, 100)
	copy(suf, cmd.Value.Summary)
	s.Suffix = string(suf)

	created, err := issue.Create(cmd.Value)
	if err != nil {
		return err
	}
	s.FinalMSG = fmt.Sprintf("%v  %v \n", prompt.IconClear, issue.Label(*created))
	return nil
}

func (cmd *CreateCommand) Response() error {
	return nil
}

func Create(option *CreateOption) error {
	buf, err := ioutil.ReadFile(option.FilePath)
	if err != nil {
		return err
	}

	values := make([]*issue.CreateValue, 0)
	err = yaml.Unmarshal(buf, &values)
	if err != nil {
		return err
	}

	for _, v := range values {
		err := prompt.Progress(&CreateCommand{
			Value: v,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
