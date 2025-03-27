package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/mend"
	"github.com/bbfh-dev/mend/mend/settings"
	"github.com/bbfh-dev/parsex/parsex"
)

const FLAG_INPUT = "input"
const FLAG_USE_TABS = "tabs"
const FLAG_INDENT = "indent"
const FLAG_DECOMMENT = "decomment"

var Version string

var CLI = parsex.New(
	"mend",
	"Mend is a simple HTML template processor designed to, but not limited to be used to generate static websites.",
	Program,
).
	SetVersion(Version).
	AddOptions(
		parsex.ParamOption{
			Name:     FLAG_INPUT,
			Keywords: []string{FLAG_INPUT, "i"},
			Desc:     "Set global input parameters",
			Check:    parsex.ValidJSON,
			Optional: true,
		},
		parsex.ParamOption{
			Name:     FLAG_INDENT,
			Keywords: []string{FLAG_INDENT},
			Desc:     "Set amount of spaces to indent with. Gets ignored if --tabs is used",
			Check:    parsex.ValidInt,
			Optional: true,
		},
		parsex.FlagOption{
			Name:     FLAG_USE_TABS,
			Keywords: []string{FLAG_USE_TABS, "t"},
			Desc:     "Use tabs instead of spaces",
		},
		parsex.FlagOption{
			Name:     FLAG_DECOMMENT,
			Keywords: []string{FLAG_DECOMMENT},
			Desc:     "Strips away any comments",
		},
	).
	AddArguments("html files...").
	Build()

func Program(input parsex.Input) error {
	if len(input.Args()) == 0 {
		return nil
	}

	var err error
	for _, filename := range input.Args() {
		filename, err = filepath.Abs(filename)
		if err != nil {
			return err
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return err
		}

		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()

		settings.KeepComments = !input.Has(FLAG_DECOMMENT)
		if input.Has(FLAG_INDENT) {
			settings.IndentWith = strings.Repeat(" ", input.Int(FLAG_INDENT))
		}
		if input.Has(FLAG_USE_TABS) {
			settings.IndentWith = "\t"
		}
		params := "{}"
		if input.Has(FLAG_INPUT) {
			params = input.String(FLAG_INPUT)
		}
		settings.GlobalParams = params

		template := mend.NewTemplate(filename, params)
		err = template.Parse(file)
		if err != nil {
			return err
		}

		template.Root.Render(os.Stdout, 0)
	}

	return nil
}

func main() {
	CLI.FromOSArgs().RunAndExit()
}
