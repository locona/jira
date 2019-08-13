package prompt

import (
	"time"

	"github.com/briandowns/spinner"
)

type Loader interface {
	Request(*spinner.Spinner) error
	Response() error
}

func Loading(loader Loader) error {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond) // Build our new spinner
	s.Start()
	s.Prefix = " loading ...  "
	s.Color("magenta")
	err := loader.Request(s)
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
