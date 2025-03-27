package mend

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

const bracketOpen = "{{"
const bracketClose = "}}"

type modifier func(original string) string

type Expression struct {
	Variable   string
	DataSource string
	Modifiers  []modifier
	Fallback   string
	IsNumber   bool
	notFound   bool
}

func (expression *Expression) String() string {
	result := gjson.Get(expression.DataSource, expression.Variable)
	if !result.Exists() {
		expression.notFound = true
		return expression.Fallback
	}
	value := result.String()

	for _, modify := range expression.Modifiers {
		value = modify(value)
	}

	return value
}

func (expression *Expression) Err() error {
	if expression.notFound && expression.Fallback == "" {
		return fmt.Errorf("undefined parameter %q", expression.Variable)
	}
	return nil
}

func ParseForExpressions(source string, text string) (string, error) {
	var builder strings.Builder
	i := 0

	for {
		start := strings.Index(text[i:], bracketOpen)
		if start == -1 {
			// No more expressions; write the remaining text.
			builder.WriteString(text[i:])
			break
		}

		builder.WriteString(text[i : i+start])
		i += start + len(bracketOpen)

		// Find the closing bracket.
		end := strings.Index(text[i:], bracketClose)
		if end == -1 {
			// Unmatched bracket; treat the remainder as plain text.
			builder.WriteString(bracketOpen)
			builder.WriteString(text[i:])
			break
		}

		// Evaluate the expression inside the brackets.
		expr, err := ComputeExpression(source, strings.TrimSpace(text[i:i+end]))
		if err != nil {
			return builder.String(), err
		}
		builder.WriteString(expr)
		i += end + len(bracketClose)
	}

	return builder.String(), nil
}
