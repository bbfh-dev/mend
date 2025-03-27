package mend

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/bbfh-dev/mend/mend/settings"
	"github.com/iancoleman/strcase"
)

const tokenGlobal = "^"
const tokenVar = "."
const tokenModifier = "-"

func ComputeExpression(source string, str string) (string, error) {
	expression := &Expression{
		Variable:   "",
		DataSource: source,
		Modifiers:  []modifier{},
		Fallback:   "",
		IsNumber:   false,
		notFound:   false,
	}

	tokens := strings.Fields(str)
	for i, token := range tokens {
		if strings.HasPrefix(token, "@") {
			// Why are we even computing an unparsed expression?
			return "", nil
		}

		if _, err := strconv.Atoi(token); err == nil {
			expression.Variable = token
			expression.IsNumber = true
			continue
		}

		if strings.HasPrefix(token, tokenGlobal) {
			expression.DataSource = settings.GlobalParams
			token = token[1:]
		}

		if strings.HasPrefix(token, tokenVar) {
			if expression.Variable != "" {
				return "", fmt.Errorf(
					"expression variable is set to %q but attempting to change it to %q",
					expression.Variable,
					token[1:],
				)
			}
			expression.Variable = token[1:]
			continue
		}

		if strings.HasPrefix(token, tokenModifier) {
			switch token[1:] {

			case "capitalize":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					if len(original) == 0 {
						return ""
					}
					return strings.ToUpper(original[:1]) + original[1:]
				})

			case "invert":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strings.Map(func(char rune) rune {
						switch {
						case unicode.IsLower(char):
							return unicode.ToUpper(char)
						case unicode.IsUpper(char):
							return unicode.ToLower(char)
						}
						return char
					}, original)
				})

			case "quote":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return fmt.Sprintf("%q", original)
				})

			case "get-length":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return fmt.Sprintf("%d", len(original))
				})

			case "get-lines":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return fmt.Sprintf("%d", strings.Count(original, "\n"))
				})

			case "get-fields":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return fmt.Sprintf("%d", len(strings.Fields(original)))
				})

			case "to-upper":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strings.ToUpper(original)
				})

			case "to-lower":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strings.ToLower(original)
				})

			case "to-snake-case":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strcase.ToSnake(original)
				})

			case "to-camel-case":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strcase.ToLowerCamel(original)
				})

			case "to-pascal-case":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strcase.ToCamel(original)
				})

			case "to-kebab-case":
				expression.Modifiers = append(expression.Modifiers, func(original string) string {
					return strcase.ToKebab(original)
				})

			default:
				return "", fmt.Errorf("unknown modifier %q", token)

			}
			continue
		}

		if token == "==" || token == "!=" {
			if len(tokens[i+1:]) == 0 {
				return "", fmt.Errorf(
					"expecting a value to the right of the equation of %q",
					expression.Variable,
				)
			}
			compare := expression.Variable
			if !expression.IsNumber {
				compare = expression.String()
			}

			var result = compare == tokens[i+1]
			if token == "!=" {
				result = !result
			}
			return fmt.Sprintf("%v", result), nil
		}

		if token == "||" {
			if len(tokens[i+1:]) == 0 {
				expression.Fallback = " "
			} else {
				expression.Fallback = strings.Join(tokens[i+1:], " ")
			}
			break
		}

		return "", fmt.Errorf("unexpected token %q", token)
	}

	if expression.Variable == "" {
		return "", fmt.Errorf("expression {{ %s }} does not contain any variables!", str)
	}

	return expression.String(), expression.Err()
}
