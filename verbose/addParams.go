package verbose

import (
	"io"
	"os"

	"github.com/nickwells/location.mod/location"
	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/param.mod/v6/psetter"
)

const (
	paramNameVerbose         = "verbose"
	paramNameVerboseToStdout = "verbose-to-stdout"
	paramNameVerboseToStderr = "verbose-to-stderr"

	paramNameTimings = "show-timings"

	paramGroupName = "pkg.verbose"
)

var verbose bool

var vDest io.Writer = os.Stdout

// setVDestToStderr
func setVDestToStderr(_ location.L, _ *param.ByName, _ []string) error {
	vDest = os.Stderr
	return nil
}

// setVDestToStdout
func setVDestToStdout(_ location.L, _ *param.ByName, _ []string) error {
	vDest = os.Stdout
	return nil
}

// setParamGroupDesc set the common description for the verbose param
// group. Note that it is safe to set the group description multiple times so
// long as the description is the same each time.
func setParamGroupDesc(ps *param.PSet) {
	ps.AddGroup(paramGroupName,
		"These are parameters to control the level of program output")
}

// AddParams will add the params to the given param set which control the
// behaviour of the Verbose function. This should be called before the
// ParamSet is parsed. All the parameters are marked as hidden
// (DontShowInStdUsage) so they will not appear in the default help message
// but will be visible if all parameters are shown.
func AddParams(ps *param.PSet) error {
	setParamGroupDesc(ps)

	ps.Add(paramNameVerbose,
		psetter.Bool{Value: &verbose},
		"set this parameter to get the verbose behaviour"+
			" - extra messages will be printed.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName))

	ps.Add(paramNameVerboseToStderr, psetter.Nil{},
		"set this parameter to have the verbose messages printed to the"+
			" standard error rather than standard out.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName),
		param.PostAction(setVDestToStderr))

	ps.Add(paramNameVerboseToStdout, psetter.Nil{},
		"set this parameter to have the verbose messages printed to the"+
			" standard output. This is the standard behaviour and"+
			" should only be needed if the destination has"+
			" already been changed, perhaps in a configuration file.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName),
		param.PostAction(setVDestToStdout))

	return nil
}

// AddTimingParams returns a param.PSetOptFunc that adds the common
// show-timing parameter used to set the ShowTimings field in a verbose.Stack
// struct
func AddTimingParams(stack *Stack) param.PSetOptFunc {
	return func(ps *param.PSet) error {
		setParamGroupDesc(ps)

		ps.Add(paramNameTimings,
			psetter.Bool{Value: &stack.ShowTimings},
			"report the time taken for parts of this program to complete.",
			param.GroupName(paramGroupName),
			param.Attrs(param.DontShowInStdUsage|param.CommandLineOnly),
			param.AltNames("show-timing", "show-time", "show-times"),
		)

		return nil
	}
}
