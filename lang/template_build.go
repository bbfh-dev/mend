package lang

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/expressions"
	"github.com/bbfh-dev/mend/lang/printer"
	"github.com/bbfh-dev/mend/lang/templating"
	"golang.org/x/net/html"
)

func (template *Template) Build(reader io.Reader) error {
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("(%s): %w", template.Cursor(), err)
		}

		template.thisToken = tokenizer.Token()
		template.thisText = strings.TrimSpace(template.thisToken.Data)
		template.thisAttrs = attrs.New(template.thisToken.Attr)
		for key, value := range template.thisAttrs.Values {
			text, err := expressions.Parse(value, template.Context.Compute)
			if err != nil {
				return fmt.Errorf("%s (expression): %w", template.Cursor(), err)
			}
			template.thisAttrs.Values[key] = text
		}

		if tokenType == html.TextToken {
			template.thisLineIndex += strings.Count(template.thisText, "\n")
		}

		if err := template.buildToken(tokenType); err != nil {
			return fmt.Errorf("(%s): %w", template.Cursor(), err)
		}
	}
}

func (template *Template) buildToken(tokenType html.TokenType) error {
	switch tokenType {

	case html.DoctypeToken:
		template.Pivot().Append(templating.NewDoctype(
			template.thisIndent,
			template.thisText,
		))

	case html.CommentToken:
		if printer.StripComments {
			break
		}
		template.Pivot().Append(templating.NewComment(
			template.thisIndent,
			template.thisText,
		))

	case html.TextToken:
		if len(template.thisText) == 0 {
			break
		}
		text, err := expressions.Parse(template.thisText, template.Context.Compute)
		if err != nil {
			return fmt.Errorf("(expression): %w", err)
		}
		template.Pivot().Append(templating.NewText(
			template.thisIndent,
			text,
		))

	case html.SelfClosingTagToken:
		switch {

		case strings.HasPrefix(template.thisText, MEND_PREFIX):
			name := strings.TrimPrefix(template.thisText, MEND_PREFIX)
			switch name {

			case "slot":
				tag := templating.NewMendSlot(template.thisIndent)
				template.Pivot().Append(tag)
				template.Slot = tag

			case "include":
				src, err := template.requireAttr(":src")
				if err != nil {
					return err
				}

				branch, err := template.BranchOut(filepath.Join(template.Dir, src))
				if err != nil {
					return err
				}

				template.Pivot().Append(branch.Root())

			default:
				return fmt.Errorf("unknown tag <%s%s />", MEND_PREFIX, name)
			}

		case strings.HasPrefix(template.thisText, PKG_PREFIX):
			name := strings.TrimPrefix(template.thisText, PKG_PREFIX)
			location, err := template.locateTemplate(name)
			if err != nil {
				return err
			}

			branch, err := template.BranchOut(location)
			if err != nil {
				return err
			}

			template.Pivot().Append(branch.Root())

		default:
			template.Pivot().Append(templating.NewSelfClosing(
				template.thisIndent,
				template.thisText,
				template.thisAttrs,
			))
		}

	case html.StartTagToken:
		switch {

		case strings.HasPrefix(template.thisText, MEND_PREFIX):
			name := strings.TrimPrefix(template.thisText, MEND_PREFIX)
			switch name {

			case "slot":
				tag := templating.NewMendSlot(template.thisIndent)
				template.EnterPivot(tag)
				template.Slot = tag

			case "if":

			case "extend":
				src, err := template.requireAttr(":src")
				if err != nil {
					return err
				}

				branch, err := template.BranchOut(filepath.Join(template.Dir, src))
				if err != nil {
					return err
				}

				template.EnterPivot(templating.NewMendExtend(
					template.thisIndent,
					branch.Root(),
					branch.Slot,
				))

			default:
				return fmt.Errorf("unknown tag <%s%s>", MEND_PREFIX, name)
			}

		case strings.HasPrefix(template.thisText, PKG_PREFIX):
			name := strings.TrimPrefix(template.thisText, PKG_PREFIX)
			location, err := template.locateTemplate(name)
			if err != nil {
				return err
			}

			branch, err := template.BranchOut(location)
			if err != nil {
				return err
			}

			template.EnterPivot(templating.NewMendExtend(
				template.thisIndent,
				branch.Root(),
				branch.Slot,
			))

		case template.thisText == "html":
			template.EnterPivot(
				templating.NewDefaultRoot(
					template.thisIndent,
					template.thisText,
					template.thisAttrs,
				),
			)

		default:
			template.EnterPivot(
				templating.NewDefault(
					template.thisIndent,
					template.thisText,
					template.thisAttrs,
				),
			)
		}

	case html.EndTagToken:
		switch tag := template.Pivot().(type) {
		case *templating.MendExtendTag:
			if tag.Slot == nil {
				fmt.Fprintf(
					os.Stderr,
					"WARN: (%s) couldn't find <mend:slot> block inside of extended file. Skipping body\n",
					template.Cursor(),
				)
			} else {
				tag.Slot.SetChildren(tag.Children)
				tag.Slot.Shift(template.thisIndent)
			}
		}
		template.ExitPivot()
	}

	return nil
}
