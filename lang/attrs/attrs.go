package attrs

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/printer"
	"golang.org/x/net/html"
)

type Attributes struct {
	order  []string
	Values map[string]string
}

func New(sourceAttrs []html.Attribute) *Attributes {
	attrs := &Attributes{
		order:  []string{},
		Values: map[string]string{},
	}

	for _, attr := range sourceAttrs {
		attrs.order = append(attrs.order, attr.Key)
		attrs.Values[attr.Key] = attr.Val
	}

	return attrs.Sort()
}

func (attrs *Attributes) Render(out printer.Writer) {
	for _, key := range attrs.order {
		out.WriteString(" ")
		attrs.renderKey(out, key)
	}
}

func (attrs *Attributes) renderKey(out printer.Writer, key string) {
	if len(attrs.Values[key]) == 0 {
		out.WriteString(key)
		return
	}

	fmt.Fprintf(out, "%s=%q", key, attrs.Values[key])
}
