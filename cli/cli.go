package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/mend"
	"github.com/bbfh-dev/mend/mend/settings"
)

var Options struct {
	Input         string `alt:"i" desc:"Set global input parameters"`
	Indent        int    `desc:"Set amount of spaces to indent with. Gets ignored if --tabs is used"`
	Tabs          bool   `alt:"t" desc:"Use tabs for indentation"`
	StripComments bool   `desc:"Strips away any comments"`
}

func Main(args []string) error {
	if len(args) == 0 {
		return nil
	}

	settings.KeepComments = !Options.StripComments
	if Options.Tabs {
		settings.IndentWith = "\t"
	} else if Options.Indent != 0 {
		settings.IndentWith = strings.Repeat(" ", Options.Indent)
	}

	var err error
	for _, filename := range args {
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
		if Options.Input != "" {
			params = Options.Input
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
