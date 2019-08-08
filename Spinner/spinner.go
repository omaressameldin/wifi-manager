package spinner

import (
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
)

type Spinner struct {
	Shape   int
	spinner *spinner.Spinner
}

const SPAN time.Duration = 100
const DELAY time.Duration = 200 * time.Millisecond

func (s *Spinner) StartSpinner(text string) {
	s.spinner = spinner.New(spinner.CharSets[s.Shape], SPAN*time.Millisecond)
	s.spinner.Prefix = fmt.Sprintf("%v ", text)

	s.spinner.Start()
}

func (s *Spinner) StartSpinnerWithDelay(text string) {
	time.Sleep(DELAY)
	s.StartSpinner(text)
}

func (s *Spinner) StopSpinner() {
	if s.spinner == nil {
		log.Fatalf("need to start spinner before stopping!")
	}
	s.spinner.Stop()
}
