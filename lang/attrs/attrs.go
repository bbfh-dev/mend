package attrs

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/mend/lang/printer"
	"golang.org/x/net/html"
)

// XML attributes in a sorted manner
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

// NOTE: It prepends " " (space) to the output
func (attrs *Attributes) Render(out printer.Writer) {
	for _, key := range attrs.order {
		out.WriteString(" ")
		attrs.renderKey(out, key)
	}
}

// Overrides or saves new attribute. Use %s to format the original attribute
func (attrs *Attributes) OverrideAttr(key string, value string) {
	original, ok := attrs.Values[key]
	if !ok {
		attrs.order = append(attrs.order, key)
		attrs.Values[key] = value
		attrs.Sort()
		return
	}

	if strings.Contains(value, "%s") {
		attrs.Values[key] = fmt.Sprintf(value, original)
	} else {
		attrs.Values[key] = value
	}
}

func (attrs *Attributes) renderKey(out printer.Writer, key string) {
	if len(attrs.Values[key]) == 0 {
		out.WriteString(key)
		return
	}

	fmt.Fprintf(out, "%s=%q", key, attrs.Values[key])
}
