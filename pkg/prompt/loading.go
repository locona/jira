package prompt

import (
	"time"

	"github.com/briandowns/spinner"
)

type Loader interface {
	Request() error
	Response() error
}

func Loading(loader Loader) error {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond) // Build our new spinner
	s.Prefix = " loading ...  "
	s.Color("magenta")
	s.Start()
	err := loader.Request()
	time.Sleep(1 * time.Second)
	s.Stop()

	if err != nil {
		return err
	}

	err = loader.Response()
	if err != nil {
		return err
	}
	return err
}
