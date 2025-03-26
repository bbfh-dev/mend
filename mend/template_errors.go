package mend

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/mend.html/mend/tags"
)

func (template *Template) errUnknownTag() error {
	var help strings.Builder
	for _, tag := range tags.AllTags {
		fmt.Fprintf(&help, "<%s> ", tag)
	}

	return fmt.Errorf(
		"unknown custom tag <%s>.\nPossible tags are: %s",
		template.currentText,
		help.String(),
	)
}

func (template *Template) errInternal(err error) error {
	return fmt.Errorf("(internal) %w", err)
}

func (template *Template) errBranch(err error) error {
	return fmt.Errorf("-> %w", err)
}

func (template *Template) errMissingAttribute(attr string) error {
	return fmt.Errorf(
		"tag <%s> is missing attribute %q",
		template.currentText,
		attr,
	)
}

func (template *Template) errUndefinedParam(param string) error {
	return fmt.Errorf("undefined parameter %q", param)
}
