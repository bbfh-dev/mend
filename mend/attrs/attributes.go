package attrs

import (
	"golang.org/x/net/html"
)

// Sorted HTML tag attributes key="value"
type Attributes struct {
	order  []string
	values map[string]string
}

func New(fromAttrs []html.Attribute) Attributes {
	out := Attributes{
		order:  []string{},
		values: map[string]string{},
	}

	for _, attr := range fromAttrs {
		out.order = append(out.order, attr.Key)
		out.values[attr.Key] = attr.Val
	}

	return out.sort()
}

func (attrs Attributes) IsEmpty() bool {
	return len(attrs.order) == 0
}

// func (attrs Attributes) ParamKeys() []string {
// 	out := make([]string, 0, len(attrs.order))
// 	for _, key := range attrs.order {
// 		if strings.HasPrefix(key, ":") {
// 			out = append(out, key[1:])
// 		}
// 	}
// 	return out
// }
//
// func (attrs Attributes) ParseExpressions(fn func(string) string) Attributes {
// 	for key, value := range attrs.values {
// 		attrs.values[key] = fn(value)
// 	}
// 	return attrs
// }

func (attrs Attributes) Get(key string) string {
	return attrs.values[key]
}

func (attrs Attributes) Contains(key string) bool {
	_, ok := attrs.values[key]
	return ok
}

func (attrs Attributes) GetOrFallback(key string, fallback string) string {
	value, ok := attrs.values[key]
	if ok {
		return value
	}
	return fallback
}
