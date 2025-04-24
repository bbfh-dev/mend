package context

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

var GlobalContext *Context

type Context struct {
	Values map[string]any
}

func New() *Context {
	return &Context{
		Values: map[string]any{},
	}
}

func (ctx *Context) Query(path []string) (string, error) {
	if len(path) == 0 {
		return ctx.String(), nil
	}

	value, ok := ctx.Values[path[0]]
	if !ok {
		return "", fmt.Errorf("undefined property %q in %q", path[0], ctx.String())
	}

	switch value := value.(type) {
	case *Context:
		return value.Query(path[1:])
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

func (ctx *Context) Set(path []string, newValue string) {
	switch len(path) {
	case 0:
		return
	case 1:
		ctx.Values[path[0]] = newValue
		return
	}

	value, ok := ctx.Values[path[0]]
	if !ok {
		value = New()
		ctx.Values[path[0]] = value
	}

	switch value := value.(type) {
	case *Context:
		value.Set(path[1:], newValue)
	case string:
		ctx.Values[path[0]] = newValue
	}
}

func (ctx *Context) String() string {
	var builder strings.Builder

	for key, value := range ctx.Values {
		switch value := value.(type) {
		case string:
			fmt.Fprintf(&builder, "%s='%s' ", key, value)
		default:
			fmt.Fprintf(&builder, "%s=%s ", key, value)
		}
	}

	str := builder.String()
	if len(str) == 0 {
		return "{}"
	}
	return "{" + str[:len(str)-1] + "}"
}

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
