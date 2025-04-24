package expressions

import "strings"

const bracketOpen = "[["
const bracketClose = "]]"

func Parse(text string, callback func(string) (string, error)) (string, error) {
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
		expr, err := callback(strings.TrimSpace(text[i : i+end]))
		if err != nil {
			return builder.String(), err
		}
		builder.WriteString(expr)
		i += end + len(bracketClose)
	}

	return builder.String(), nil
}
