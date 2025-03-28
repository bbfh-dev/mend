package attrs

import (
	"strings"

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

func (attrs Attributes) ParamKeys() map[string]string {
	out := map[string]string{}
	for _, key := range attrs.order {
		if strings.HasPrefix(key, ":") {
			out[key[1:]] = attrs.values[key]
		}
	}
	return out
}

func (attrs Attributes) InheritAttributes() Attributes {
	out := New([]html.Attribute{})
	for _, key := range attrs.order {
		if strings.HasPrefix(key, "@") {
			out.order = append(out.order, key[1:])
			out.values[key[1:]] = attrs.values[key]
		}
	}
	return out
}

func (attrs Attributes) Merge(overwrite Attributes) Attributes {
	for _, key := range overwrite.order {
		_, ok := attrs.values[key]
		if !ok {
			attrs.order = append(attrs.order, key)
		}
		attrs.values[key] = overwrite.values[key]
	}
	return attrs
}

func (attrs Attributes) ParseExpressions(
	source string,
	fn func(string, string) (string, error),
) (Attributes, error) {
	for key, value := range attrs.values {
		newValue, err := fn(source, value)
		if err != nil {
			return attrs, err
		}
		attrs.values[key] = newValue
	}
	return attrs, nil
}

func (attrs Attributes) ReplaceText(text string, with string) Attributes {
	clone := Attributes{
		order:  attrs.order,
		values: map[string]string{},
	}
	for key, value := range attrs.values {
		clone.values[key] = strings.ReplaceAll(value, text, with)
	}
	return clone
}

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
