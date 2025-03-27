package mend

import (
	"github.com/bbfh-dev/mend/mend/assert"
	"github.com/bbfh-dev/mend/mend/attrs"
	"github.com/bbfh-dev/mend/mend/tags"
	"golang.org/x/net/html"
)

const PREFIX = "mend:"
const PREFIX_LEN = len(PREFIX)

type Template struct {
	Name   string
	Params string

	Root tags.NodeWithChildren
	Slot tags.NodeWithChildren

	currentLine  int
	currentToken html.Token
	currentText  string
	currentAttrs attrs.Attributes
	// A list of current parents from greatest to closest
	breadcrumbs []tags.NodeWithChildren
}

func NewTemplate(name string, params string) *Template {
	root := tags.NewRootNode()
	return &Template{
		Name:         name,
		Params:       params,
		Root:         root,
		Slot:         nil,
		currentLine:  1,
		currentToken: html.Token{},
		currentAttrs: attrs.Attributes{},
		currentText:  "",
		breadcrumbs:  []tags.NodeWithChildren{root},
	}
}

func (template *Template) lastBreadcrumb() tags.NodeWithChildren {
	assert.NotEmpty(template.breadcrumbs, "template must never have no breadcrumbs left")
	return template.breadcrumbs[len(template.breadcrumbs)-1]
}

func (template *Template) grandParent() tags.NodeWithChildren {
	switch len(template.breadcrumbs) {
	case 0, 1:
		return template.Root
	}
	return template.breadcrumbs[len(template.breadcrumbs)-2]
}

func (template *Template) append(nodes ...tags.Node) {
	template.lastBreadcrumb().Add(nodes...)
}

func (template *Template) appendLevel(node tags.NodeWithChildren) {
	template.append(node)
	template.breadcrumbs = append(template.breadcrumbs, node)
}
