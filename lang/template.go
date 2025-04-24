package lang

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/context"
	"github.com/bbfh-dev/mend/lang/templating"
	"golang.org/x/net/html"
)

const MEND_PREFIX = "mend:"
const PKG_PREFIX = "pkg:"

type Template struct {
	Dir     string
	Name    string
	Context *context.Context

	Breadcrumbs []templating.PairedTag
	Slot        *templating.MendSlotTag

	thisToken     html.Token
	thisText      string
	thisAttrs     *attrs.Attributes
	thisLineIndex int
	thisIndent    int
}

func New(indent int, ctx *context.Context, dir, name string) *Template {
	return &Template{
		Dir:           dir,
		Name:          name,
		Context:       ctx,
		Breadcrumbs:   []templating.PairedTag{templating.NewMendSlot(indent)},
		Slot:          nil,
		thisToken:     html.Token{},
		thisText:      "",
		thisAttrs:     nil,
		thisLineIndex: 0,
		thisIndent:    indent,
	}
}

func (template *Template) Cursor() string {
	return fmt.Sprintf("%s:%d", template.Name, template.thisLineIndex+1)
}

func (template *Template) Root() templating.PairedTag {
	return template.Breadcrumbs[0]
}

func (template *Template) Pivot() templating.PairedTag {
	return template.Breadcrumbs[len(template.Breadcrumbs)-1]
}

func (template *Template) EnterPivot(tag templating.PairedTag) {
	template.Pivot().Append(tag)
	template.Breadcrumbs = append(template.Breadcrumbs, tag)
	template.thisIndent++
}

func (template *Template) ExitPivot() {
	if len(template.Breadcrumbs) == 1 {
		return
	}
	template.Breadcrumbs = template.Breadcrumbs[:len(template.Breadcrumbs)-1]
	template.thisIndent--
}

func (template *Template) requireAttr(key string) (string, error) {
	src, ok := template.thisAttrs.Values[key]
	if !ok {
		return "", fmt.Errorf(
			"<%s> requires an `:%s=\"...\"` attribute",
			template.thisText,
			key,
		)
	}
	return src, nil
}
