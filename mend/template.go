package mend

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/bbfh-dev/mend.html/mend/attrs"
	"golang.org/x/net/html"
)

type Template struct {
	Name   string
	Params string

	currentLine  int
	currentToken html.Token
	currentAttrs attrs.Attributes
}

func NewTemplate(name string, params string) *Template {
	return &Template{
		Name:        name,
		Params:      params,
		currentLine: 1,
	}
}

func (template *Template) Parse(reader io.Reader) error {
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			return fmt.Errorf(
				"(%s) %w",
				filepath.Base(template.Name),
				tokenizer.Err(),
			)
		}

		template.currentToken = tokenizer.Token()
		template.currentAttrs = attrs.New(template.currentToken.Attr)

		err := template.Process(tokenType)
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

func (template *Template) Process(tokenType html.TokenType) error {
	return nil
}
