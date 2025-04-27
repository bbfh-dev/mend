package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/lang"
	"github.com/bbfh-dev/mend/lang/context"
	"github.com/bbfh-dev/mend/lang/printer"
)

var Options struct {
	Tabs          bool   `alt:"t" desc:"Use tabs for indentation"`
	Indent        int    `default:"4" desc:"The amount of spaces to be used for indentation (overwriten by --tabs)"`
	StripComments bool   `desc:"Strips away HTML comments from the output"`
	Input         string `desc:"Provide input to the provided files in the following format: 'attr1=value1,attr2.a.b.c=value2,...'"`
	Output        string `desc:"(Optional) output path. Use '.' to substitute the same filename (e.g. './out/.' -> './out/input.html')"`
	WorkDir       string `desc:"Overwrite the directory used for resolving <pkg:...> tags"`
}

func Main(args []string) error {
	printer.StripComments = Options.StripComments
	if Options.Tabs {
		printer.IndentString = "\t"
	} else {
		printer.IndentString = strings.Repeat(" ", Options.Indent)
	}

	lang.Cwd = Options.WorkDir
	context.GlobalContext = context.New()
	if len(Options.Input) != 0 {
		for prop := range strings.SplitSeq(Options.Input, ",") {
			pair := strings.SplitN(prop, "=", 2)
			if len(pair) != 2 {
				return errors.New("input format must be: 'attr1=value1,attr2.a.b.c=value2,...'")
			}
			key, value := pair[0], pair[1]
			context.GlobalContext.Set(strings.Split(key, "."), value)
		}
	}

	for _, filename := range args {
		dir := filepath.Dir(filename)
		base := filepath.Base(filename)
		filename, _ := filepath.Abs(filename)

		context.GlobalContext.Set([]string{"@file"}, filename)
		if Options.WorkDir == "" {
			lang.Cwd = dir
		}

		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()

		template := lang.New(0, context.GlobalContext, dir, base)
		if err := template.Build(file); err != nil {
			return err
		}

		out := os.Stdout
		if Options.Output != "" {
			if strings.HasSuffix(Options.Output, ".") {
				Options.Output = strings.TrimSuffix(Options.Output, ".") + base
			}

			if err := os.MkdirAll(filepath.Dir(Options.Output), os.ModePerm); err != nil {
				return err
			}

			out, err = os.OpenFile(Options.Output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
			if err != nil {
				return err
			}
		}

		template.Root().Render(out, -1)
	}

	return nil
}
