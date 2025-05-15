package context

import (
	"strings"

	"github.com/bbfh-dev/mend/lang/attrs"
)

func IsContextKey(key string) bool {
	return strings.HasPrefix(key, ":")
}

func ParseAttrs(attrs *attrs.Attributes) *Context {
	ctx := New()

	for key, attr := range attrs.Values {
		if !IsContextKey(key) {
			continue
		}
		key = key[1:]
		value := strings.TrimSpace(attr)

		if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
			ctx.Values[key] = parseDict(value)
			continue
		}

		ctx.Set([]string{key}, strings.TrimPrefix(value, "&"))
	}

	return ctx
}

func parseDict(str string) *Context {
	dict := New()
	if len(str) == 2 {
		return dict
	}

	content := strings.TrimSpace(str[1 : len(str)-1])
	for len(content) > 0 {
		// key up to '='
		eq := strings.Index(content, "=")
		if eq < 0 {
			break
		}
		key := content[:eq]
		content = content[eq+1:]
		if len(content) == 0 {
			break
		}

		switch {
		// nested dict
		case content[0] == '{':
			depth := 1
			i := 1
			for ; i < len(content) && depth > 0; i++ {
				switch content[i] {
				case '{':
					depth++
				case '}':
					depth--
				}
			}
			sub := content[:i]
			dict.Values[key] = parseDict(sub)
			content = strings.TrimLeft(content[i:], " ")

		// quoted string
		case content[0] == '\'':
			i := 1
			for ; i < len(content) && content[i] != '\''; i++ {
			}
			dict.Values[key] = content[1:i]
			content = strings.TrimLeft(content[i+1:], " ")

		// bare word
		default:
			i := strings.Index(content, " ")
			if i < 0 {
				dict.Values[key] = content
				content = ""
			} else {
				dict.Values[key] = content[:i]
				content = strings.TrimLeft(content[i+1:], " ")
			}
		}
	}

	return dict
}
