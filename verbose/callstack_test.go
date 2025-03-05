package verbose

import (
	"bytes"
	"os"
	"regexp"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
	"golang.org/x/exp/slices"
)

func TestStack(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		tags     []string
		expTagRE *regexp.Regexp
	}{
		{
			ID:   testhelper.MkID("..."),
			tags: []string{"a", "b", "c"},
			expTagRE: regexp.MustCompile("a[.]{29}: Start\n" +
				"|    b[.]{24}: Start\n" +
				"|    |    c[.]{19}: Start\n" +
				"|    |    c[.]{19}:[ .0-9] msecs\n" +
				"|    b[.]{24}:[ .0-9] msecs\n" +
				"a[.]{29}:[ .0-9] msecs\n"),
		},
	}

	defer func() { vDest = os.Stdin }()

	for _, tc := range testCases {
		stack := &Stack{ShowTimings: true}
		buf := new(bytes.Buffer)
		vDest = buf
		funcStack := []func(){}

		for _, tag := range tc.tags {
			funcStack = append(funcStack, stack.Start(tag, "msg"))
		}

		slices.Reverse[[]func(), func()](funcStack)

		for _, f := range funcStack {
			f()
		}

		if !tc.expTagRE.MatchString(buf.String()) {
			t.Logf("%s\n", tc.IDStr())
			t.Errorf("\t: bad output: %q", buf.String())
		}
	}
}
