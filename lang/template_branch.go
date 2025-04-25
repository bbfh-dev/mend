package lang

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/lang/context"
)

func (template *Template) BranchOut(location string) (*Template, error) {
	file, err := os.OpenFile(location, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	branch := New(
		template.thisIndent,
		context.ParseAttrs(template.thisToken.Attr),
		filepath.Dir(location),
		filepath.Base(location),
	)
	if err := branch.Build(file); err != nil {
		return nil, err
	}

	for key, attr := range template.thisAttrs.Values {
		if strings.HasPrefix(key, ":") {
			continue
		}
		branch.Root().OverrideAttr(key, attr)
	}

	return branch, nil
}

func (template *Template) Find(name string) (string, bool) {
	var result string

	filepath.WalkDir(template.Dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		if entry.Name() == name {
			result = path
			return filepath.SkipAll
		}

		return nil
	})

	return result, result != ""
}

func (template *Template) locateTemplate(tag string) (string, error) {
	location, found := template.Find(tag + ".html")
	if !found {
		abs, _ := filepath.Abs(template.Dir)
		return "", fmt.Errorf(
			"Can't find template <%s%s> in %s/*",
			PKG_PREFIX,
			tag,
			abs,
		)
	}
	return location, nil
}
