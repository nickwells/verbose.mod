package verbose

import (
	"fmt"
	"strings"
	"time"

	"github.com/nickwells/timer.mod/timer"
)

const maxStackWidth = 30

// Stack used in conjunction with the timer and verbose packages this
// will print out how long a function took to run
type Stack struct {
	ShowTimings bool
	stack       []string
}

// Start prints the Start message, starts a timer and returns the function to
// be called at the end. The tag is the name of the function or code section
// that is being timed, the msg will only be printed in verbose mode.
func (s *Stack) Start(tag, msg string) func() {
	s.stack = append(s.stack, tag)
	if verbose {
		fmt.Fprintln(vDest, s.Tag(), msg)
	} else if s.ShowTimings {
		fmt.Fprintln(vDest, s.Tag(), "Start")
	} else {
		return func() { s.popStack() }
	}

	return timer.Start(tag, s)
}

// Tag returns a stacked tag reflecting the current stack depth and
// right-filled.
func (s *Stack) Tag() string {
	if len(s.stack) == 0 {
		return "[empty]"
	}

	t := strings.Repeat("|    ", len(s.stack)-1) +
		s.stack[len(s.stack)-1]

	if len(t) < maxStackWidth {
		t += strings.Repeat(".", maxStackWidth-len(t))
	}

	t += ":"

	return t
}

// popStack removes the last stack entry
func (s *Stack) popStack() {
	if len(s.stack) == 0 {
		return
	}

	s.stack = s.stack[:len(s.stack)-1]
}

// Act satisfies the action function interface for a timer. It prints out the
// tag and the duration in milliseconds if the program is in verbose mode
func (s *Stack) Act(_ string, d time.Duration) {
	if verbose || s.ShowTimings {
		fmt.Fprintf(vDest, "%s%12.3f msecs\n",
			s.Tag(), float64(d/time.Millisecond))
	}

	s.popStack()
}
