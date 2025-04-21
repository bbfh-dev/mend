package mend

import (
	"io/fs"
	"path/filepath"
)

func (template *Template) Find(name string) (string, bool) {
	var result string

	filepath.WalkDir(
		filepath.Dir(template.Name),
		func(path string, entry fs.DirEntry, err error) error {
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
		},
	)

	return result, result != ""
}
