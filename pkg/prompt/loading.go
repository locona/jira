package prompt

import (
	"time"

	"github.com/briandowns/spinner"
)

type ProgressIF interface {
	Request(*spinner.Spinner) error
	Response() error
}

func Progress(progress ProgressIF) error {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond) // Build our new spinner
	s.Color("magenta")
	s.Start()
	err := progress.Request(s)
	time.Sleep(1 * time.Second)
	s.Stop()

	if err != nil {
		return err
	}

	err = progress.Response()
	if err != nil {
		return err
	}
	return err
}
