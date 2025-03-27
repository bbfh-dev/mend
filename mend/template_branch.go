package mend

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/mend/std"
)

func (template *Template) branchOut() (*Template, error) {
	if !template.currentAttrs.Contains("src") {
		return nil, template.errMissingAttribute("src")
	}
	src := template.currentAttrs.Get("src")
	var file fs.File
	var err error

	if strings.HasPrefix(src, std.PREFIX) {
		file, err = std.Open(src[std.PREFIX_LEN:])
	} else {
		src = filepath.Join(filepath.Dir(template.Name), src)
		file, err = os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	}

	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := json.Marshal(template.currentAttrs.ParamKeys())
	if err != nil {
		return nil, template.errInternal(err)
	}

	branch := NewTemplate(src, string(data))
	err = branch.Parse(file)
	if err != nil {
		return nil, template.errBranch(err)
	}

	return branch, nil
}
