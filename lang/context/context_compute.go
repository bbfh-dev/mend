package context

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

func (ctx *Context) Compute(expression string) (string, error) {
	fields := getFields(expression)
	switch len(fields) {
	case 0:
		return "", nil
	case 1:
		return ctx.queryPath(fields[0])
	}

	variable, err := ctx.queryPath(fields[0])
	if err != nil && !slices.Contains(fields, "||") {
		return "", err
	}

	if len(fields) < 3 {
		return "", fmt.Errorf("expression requires a value to the right of %q", fields[1])
	}

	var result string
	switch fields[1] {
	case "==":
		result = fmt.Sprintf("%v", variable == fields[2])
	case "!=":
		result = fmt.Sprintf("%v", variable != fields[2])
	case "has":
		result = fmt.Sprintf("%v", strings.Contains(variable, fields[2]))
	case "lacks":
		result = fmt.Sprintf("%v", !strings.Contains(variable, fields[2]))
	case "||":
		if err != nil {
			result = fields[2]
			break
		}
		result = variable
	default:
		return "", fmt.Errorf("unknown operation %q", fields[1])
	}

	if len(fields) < 5 {
		return result, nil
	}

	switch fields[3] {
	case "||":
		if err != nil {
			result = fields[4]
			break
		}
	default:
		return "", fmt.Errorf(
			"unsupported after-operation %q, only || is supported for chaining",
			fields[3],
		)
	}

	return result, nil
}

func (ctx *Context) queryPath(fieldPath string) (result string, err error) {
	field := strings.TrimSuffix(fieldPath, "?")
	segments := strings.Split(field, ".")[1:]
	path := make([]string, 0, len(segments))
	operations := make([]string, 0, len(segments))

	for _, segment := range segments {
		if strings.HasSuffix(segment, "()") {
			operations = append(operations, segment[:len(segment)-2])
			continue
		}
		path = append(path, segment)
	}

	if strings.HasPrefix(field, "this.") {
		result, err = ctx.Query(path)
	}
	if strings.HasPrefix(field, "root.") {
		result, err = GlobalContext.Query(path)
	}

	if err != nil {
		if strings.HasSuffix(fieldPath, "?") {
			return "", nil
		}
		return "", fmt.Errorf("%s: %w", field, err)
	}

	for _, operation := range operations {
		switch operation {
		case "to_lower":
			result = strings.ToLower(result)
		case "to_upper":
			result = strings.ToUpper(result)
		case "to_pascal_case":
			result = strcase.ToCamel(result)
		case "to_camel_case":
			result = strcase.ToLowerCamel(result)
		case "to_snake_case":
			result = strcase.ToSnake(result)
		case "to_kebab_case":
			result = strcase.ToKebab(result)
		case "capitalize":
			result = strings.ToUpper(result[:1]) + result[1:]
		case "length":
			result = fmt.Sprintf("%d", len(result))
		case "filename":
			result = filepath.Base(result)
		case "dir":
			result = filepath.Dir(result)
		case "extension":
			result = filepath.Ext(result)
		case "trim_extension":
			result = strings.TrimSuffix(result, filepath.Ext(result))
		default:
			return "", fmt.Errorf("%s: unknown operation %q on %q", field, operation+"()", result)
		}
	}

	return result, nil
}

func getFields(text string) (out []string) {
	var (
		current   []rune
		inQuote   bool
		quoteChar rune
	)

	flush := func() {
		if len(current) > 0 {
			out = append(out, string(current))
			current = current[:0]
		}
	}

	for _, char := range text {
		switch {
		case (char == '"' || char == '\''):
			if inQuote {
				// closing quote?
				if char == quoteChar {
					inQuote = false
					quoteChar = 0
				} else {
					// different quote inside a quote: treat literally
					current = append(current, char)
				}
			} else {
				inQuote = true
				quoteChar = char
			}

		case unicode.IsSpace(char):
			if inQuote {
				current = append(current, char)
			} else {
				flush()
			}

		default:
			current = append(current, char)
		}
	}

	flush()
	return
}
