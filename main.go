package main

import (
	"os"
	"path/filepath"

	"github.com/bbfh-dev/mend.html/mend"
	"github.com/bbfh-dev/parsex/parsex"
)

const FLAG_INPUT = "input"
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

		params := "{}"
		if input.Has(FLAG_INPUT) {
			params = input.String(FLAG_INPUT)
		}
		template := mend.NewTemplate(filename, params)
		err = template.Parse(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	CLI.FromOSArgs().RunAndExit()
}
