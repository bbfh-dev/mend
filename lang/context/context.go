package context

import (
	"fmt"
	"strings"
)

// The 'root.' context
var GlobalContext *Context

// Template's context that's accessed using `this.`
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
