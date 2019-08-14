package project

import (
	"github.com/3-shake/jira/pkg/project"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewCommandNamespace() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ns",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Namespace()
		},
	}

	return cmd
}

type Item struct {
	Name      string
	IsCurrent bool
}

var templates = &promptui.SelectTemplates{
	Active:   "{{ .Name | cyan }}",
	Inactive: "{{ if .IsCurrent }} {{ .Name | cyan }} {{else}} {{ .Name | white }} {{end}}",
	Selected: "{{ .Name | red }}",
}

func Namespace() error {
	currentProject, err := project.Current()
	if err != nil {
		return err
	}
	list, err := project.List()
	if err != nil {
		return err
	}

	items := make([]Item, 0)
	for _, v := range *list {
		isCurrent := false
		if v.Key == currentProject {
			isCurrent = true
		}
		items = append(items, Item{
			Name:      v.Key,
			IsCurrent: isCurrent,
		})
	}

	prompt := promptui.Select{
		Label:     "Project Name",
		Items:     items,
		Templates: templates,
		Size:      10,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return err
	}

	err = project.Store(items[i].Name)
	if err != nil {
		return err
	}

	return nil
}
