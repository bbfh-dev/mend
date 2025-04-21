package mend

import (
	"encoding/json"
	"io/fs"
	"os"
)

func (template *Template) branchOut(src string) (*Template, error) {
	var file fs.File
	var err error

	file, err = os.OpenFile(src, os.O_RDONLY, os.ModePerm)

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

	branch.Root.MergeAttributes(template.currentAttrs.InheritAttributes())

	return branch, nil
}
