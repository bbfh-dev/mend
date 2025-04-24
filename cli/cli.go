package cli

import (
	"strings"

	"github.com/bbfh-dev/mend/lang/printer"
)

var Options struct {
	Tabs          bool   `alt:"t" desc:"Use tabs for indentation"`
	Indent        int    `default:"4" desc:"The amount of spaces to be used for indentation (overwriten by --tabs)"`
	StripComments bool   `desc:"Strips away HTML comments from the output"`
	Input         string `desc:"Provide input to the provided files in the following format: 'attr1=value1,attr2=value2,...'"`
}

func Main(args []string) error {
	printer.StripComments = Options.StripComments
	if Options.Tabs {
		printer.IndentString = "\t"
	} else {
		printer.IndentString = strings.Repeat(" ", Options.Indent)
	}

	return nil
}
