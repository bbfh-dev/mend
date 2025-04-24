package templating

import (
	"strings"

	"github.com/bbfh-dev/mend/lang/printer"
)

type TextTag struct {
	*BaseTag
	Text string
}

func NewText(indent int, text string) *TextTag {
	return &TextTag{
		BaseTag: NewBase(indent),
		Text:    text,
	}
}

func (tag *TextTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)

	lines := strings.Split(tag.Text, "\n")
	lastLine := len(lines) - 1
	for i, line := range lines {
		line = strings.TrimSpace(line)
		writer.WriteString(line)
		if i != lastLine {
			writer.WriteString(" ")
		}
	}
}
