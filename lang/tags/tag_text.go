package tags

import (
	"strings"

	"github.com/bbfh-dev/mend/lang/printer"
)

type TextTag struct {
	*BaseTag
	Text string
}

func NewText(text string) *TextTag {
	return &TextTag{
		BaseTag: NewBase(),
		Text:    text,
	}
}

func (tag *TextTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)

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
