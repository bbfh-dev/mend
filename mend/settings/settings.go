package settings

import (
	"io"
	"strings"
)

var GlobalParams string

var KeepComments bool

var IndentWith string = strings.Repeat(" ", 4)

func WriteIndent(output io.StringWriter, indent int) {
	output.WriteString(strings.Repeat(IndentWith, indent))
}
