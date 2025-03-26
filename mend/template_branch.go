package mend

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (template *Template) branchOut() (*Template, error) {
	if !template.currentAttrs.Contains("src") {
		return nil, template.errMissingAttribute("src")
	}
	src := filepath.Join(
		filepath.Dir(template.Name),
		template.currentAttrs.Get("src"),
	)

	file, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
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
