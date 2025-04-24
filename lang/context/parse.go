package context

import (
	"strings"

	"golang.org/x/net/html"
)

func IsContextKey(key string) bool {
	return strings.HasPrefix(key, ":")
}

func ParseAttrs(attrs []html.Attribute) *Context {
	ctx := New()

	for _, attr := range attrs {
		if !IsContextKey(attr.Key) {
			continue
		}
		key := attr.Key[1:]
		value := strings.TrimSpace(attr.Val)

		if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
			ctx.Values[key] = parseDict(value)
			continue
		}

		ctx.Set([]string{key}, value)
	}

	return ctx
}

func parseDict(str string) *Context {
	dict := New()
	if len(str) == 2 {
		return dict
	}

	return dict
}
