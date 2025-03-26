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
