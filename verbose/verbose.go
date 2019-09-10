package verbose

import "fmt"

// PrintMsgs prints the messages to the chosen destination if verbose is set
// and nothing otherwise. Each message is followed by a newline and the
// second and subsequent messages are prefixed with a tab. It is legitimate
// (but redundant) to pass no messages.
func PrintMsgs(messages ...string) {
	if !verbose {
		return
	}

	sep := ""
	for _, msg := range messages {
		fmt.Fprintln(vDest, sep, msg)
		sep = "\t"
	}
}

// Print prints the messages to the chosen destination if verbose is set and
// nothing otherwise. The messages are printed as is with no separators or
// added newlines.
func Print(messages ...string) {
	if !verbose {
		return
	}

	for _, msg := range messages {
		fmt.Fprint(vDest, msg)
	}
}

// Println prints the messages using the Print function but with an added
// newline at the end
func Println(messages ...string) {
	Print(messages...)
	Print("\n")
}

// IsOn returns true if the verbose flag has been set false otherwise
func IsOn() bool {
	return verbose
}
