package lang

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/context"
	"github.com/bbfh-dev/mend/lang/templating"
	"golang.org/x/net/html"
)

const MEND_PREFIX = ":"
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
}

func New(indent int, ctx *context.Context, dir, name string) *Template {
	return &Template{
		Dir:           dir,
		Name:          name,
		Context:       ctx,
		Breadcrumbs:   []templating.PairedTag{templating.NewPairedBase(indent)},
		Slot:          nil,
		thisToken:     html.Token{},
		thisText:      "",
		thisAttrs:     nil,
		thisLineIndex: 0,
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
}

func (template *Template) ExitPivot() {
	if len(template.Breadcrumbs) == 1 {
		return
	}
	template.Breadcrumbs = template.Breadcrumbs[:len(template.Breadcrumbs)-1]
}
