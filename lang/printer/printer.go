package printer

import (
	"io"
	"strings"
)

var IndentString string
var StripComments bool

type Writer interface {
	io.Writer
	io.StringWriter
}

func WriteIndent(writer Writer, indent int) {
	if indent < 0 {
		return
	}
	writer.WriteString(strings.Repeat(IndentString, indent))
}
