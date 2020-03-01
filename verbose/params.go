package verbose

import (
	"os"

	"github.com/nickwells/location.mod/location"
	"github.com/nickwells/param.mod/v4/param"
	"github.com/nickwells/param.mod/v4/param/psetter"
)

var verbose bool
var vDest = os.Stdout

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

// AddParams will add the params to the given param set which control the
// behaviour of the Verbose function. This should be called before the
// ParamSet is parsed. All the parameters are marked as hidden
// (DontShowInStdUsage) so they will not appear in the default help message
// but will be visible if all parameters are shown.
func AddParams(ps *param.PSet) error {
	const paramGroupName = "pkg.verbose"
	ps.AddGroup(paramGroupName,
		"These are parameters to control the level of program output")

	ps.Add("verbose", psetter.Bool{Value: &verbose},
		"set this parameter to get the verbose behaviour"+
			" - extra messages will be printed.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName))

	ps.Add("verbose-to-stderr", psetter.Nil{},
		"set this parameter to have messages printed to the"+
			" standard error rather than standard out.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName),
		param.PostAction(setVDestToStderr))

	ps.Add("verbose-to-stdout", psetter.Nil{},
		"set this parameter to have messages printed to the"+
			" standard output. This is the standard behaviour and"+
			" should only be needed if the destination has"+
			" already been changed, perhaps in a configuration file.",
		param.Attrs(param.DontShowInStdUsage),
		param.GroupName(paramGroupName),
		param.PostAction(setVDestToStdout))

	return nil
}
