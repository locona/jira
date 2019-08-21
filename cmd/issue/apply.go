package issue

import (
	"io/ioutil"
	"log"

	"github.com/3-shake/jira/pkg/issue"
	"github.com/3-shake/jira/pkg/prompt"
	"github.com/andygrunwald/go-jira"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type ApplyOption struct {
	FilePath string
}

func NewCommandApply() *cobra.Command {
	applyOption := &ApplyOption{}
	cmd := &cobra.Command{
		Use: "apply",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Apply(applyOption)
		},
	}

	cmd.Flags().StringVarP(&applyOption.FilePath, "file", "f", "", "Status.")
	return cmd
}

type ApplyCommand struct {
	Value  *issue.ApplyValue
	Result []jira.Issue
}

func (cmd *ApplyCommand) Request(s *spinner.Spinner) error {
	applied, err := issue.Apply(cmd.Value)
	if err != nil {
		return err
	}

	cmd.Result = applied
	return nil
}

func (cmd *ApplyCommand) Response() error {
	issue.ViewTable(cmd.Result)
	return nil
}

func Apply(option *ApplyOption) error {
	buf, err := ioutil.ReadFile(option.FilePath)
	if err != nil {
		return err
	}

	values := make([]*issue.ApplyValue, 0)
	err = yaml.Unmarshal(buf, &values)
	if err != nil {
		return err
	}

	for _, v := range values {
		err := prompt.Progress(&ApplyCommand{
			Value: v,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
