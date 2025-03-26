package mend

import (
	"fmt"
	"io"
	"path/filepath"
	"slices"
	"strings"

	"github.com/bbfh-dev/mend.html/mend/attrs"
	"github.com/bbfh-dev/mend.html/mend/settings"
	"github.com/bbfh-dev/mend.html/mend/tags"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html"
)

func (template *Template) Parse(reader io.Reader) error {
	tokenizer := html.NewTokenizer(reader)

loop:
	for {
		tokenType := tokenizer.Next()
		switch tokenType {

		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				break loop
			}
			return fmt.Errorf(
				"(%s) %w",
				filepath.Base(template.Name),
				tokenizer.Err(),
			)

		case html.TextToken:
			template.currentLine += strings.Count(template.currentToken.Data, "\n")

		}

		template.currentToken = tokenizer.Token()
		template.currentAttrs = attrs.New(template.currentToken.Attr)
		template.currentText = strings.TrimSpace(template.currentToken.Data)

		err := template.process(tokenType)
		if err != nil {
			return fmt.Errorf(
				"(%s:%d) %w",
				filepath.Base(template.Name),
				template.currentLine,
				err,
			)
		}
	}

	return nil
}

func (template *Template) process(tokenType html.TokenType) error {
	switch tokenType {

	case html.DoctypeToken:
		template.append(tags.NewDoctypeNode(template.currentText))

	case html.CommentToken:
		if settings.KeepComments {
			template.append(tags.NewCommentNode(template.currentText))
		}

	case html.TextToken:
		if len(template.currentText) == 0 {
			break
		}

		var builder strings.Builder
		for _, line := range strings.Split(template.currentText, "\n") {
			builder.WriteString(strings.TrimSpace(line))
			builder.WriteString(" ")
		}
		template.append(tags.NewTextNode(builder.String()))

	case html.SelfClosingTagToken:
		if !strings.HasPrefix(template.currentText, PREFIX) {
			node := tags.NewVoidNode(template.currentText, template.currentAttrs)
			template.append(node)
			break
		}
		switch template.currentText[PREFIX_LEN:] {

		case tags.TAG_INCLUDE:
			branch, err := template.branchOut()
			if err != nil {
				return err
			}
			template.append(branch.Root)

		case tags.TAG_SLOT:
			node := tags.NewRootNode()
			template.append(node)
			template.Slot = node

		default:
			return template.errUnknownTag()
		}

	case html.StartTagToken:
		if !strings.HasPrefix(template.currentText, PREFIX) {
			// Is it actually a self-closing tag with wrong syntax?
			if slices.Contains(attrs.SelfClosingTags, template.currentText) {
				return template.process(html.SelfClosingTagToken)
			}

			node := tags.NewTagNode(template.currentText, template.currentAttrs)
			template.appendLevel(node)
			break
		}
		switch template.currentText[PREFIX_LEN:] {

		case tags.TAG_EXTEND:
			branch, err := template.branchOut()
			if err != nil {
				return err
			}
			node := tags.NewCustomExtendNode()
			node.Inner.Add(branch.Root)
			node.Slot = branch.Slot
			template.appendLevel(node)

		case tags.TAG_RANGE:
			if !template.currentAttrs.Contains("for") {
				return template.errMissingAttribute("for")
			}
			variable := template.currentAttrs.Get("for")

			var result gjson.Result
			if strings.HasPrefix(variable, "^.") {
				result = gjson.Get(settings.GlobalParams, variable[2:])
			} else {
				result = gjson.Get(template.Params, variable)
			}

			if !result.Exists() {
				return template.errUndefinedParam(variable)
			}
			if !result.IsArray() {
				return fmt.Errorf(
					"parameter %q is not an array! It's set to: `%s`",
					variable,
					result.String(),
				)
			}
			node := tags.NewCustomRangeNode(variable, result)
			template.appendLevel(node)

		case tags.TAG_IF:
			node := tags.NewCustomIfNode(
				template.currentAttrs.GetOrFallback("value", "true"),
				true,
			)
			template.appendLevel(node)

		case tags.TAG_UNLESS:
			node := tags.NewCustomIfNode(
				template.currentAttrs.GetOrFallback("value", "true"),
				false,
			)
			template.appendLevel(node)

		default:
			return template.errUnknownTag()
		}

	case html.EndTagToken:
		if len(template.breadcrumbs) == 1 {
			break
		}
		switch node := template.lastBreadcrumb().(type) {

		case *tags.CustomExtendNode:
			if node.Slot != nil {
				node.Slot.Add(node.Children...)
			}

		case *tags.CustomRangeNode:
			for range node.Values.Array() {
				template.grandParent().Add(node.Children...)
			}

		}
		template.breadcrumbs = template.breadcrumbs[:len(template.breadcrumbs)-1]
	}

	return nil
}
