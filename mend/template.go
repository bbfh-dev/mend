package mend

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const PREFIX = "<!-- "
const SUFFIX = " -->"

func substring(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

type Template struct {
	writingTo string
	sections  map[string]*bytes.Buffer
}

func (template *Template) Output() string {
	var result = BASE_HTML
	for key, buffer := range template.sections {
		result = strings.Replace(
			result,
			PREFIX+"place "+key+SUFFIX,
			strings.TrimSpace(buffer.String()),
			1,
		)
	}
	return result
}

func (template *Template) Write(content string) {
	if _, ok := template.sections[template.writingTo]; !ok {
		template.sections[template.writingTo] = &bytes.Buffer{}
	}
	template.sections[template.writingTo].WriteString(content + "\n")
}

func ParseTemplate(body string) *Template {
	template := &Template{
		writingTo: "body",
		sections:  map[string]*bytes.Buffer{},
	}

	scanner := bufio.NewScanner(strings.NewReader(body))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		index := substring(line, PREFIX)
		if index == -1 {
			template.Write(line)
			continue
		}

		indexEnd := substring(line, SUFFIX)
		if indexEnd == -1 {
			panic("Mend tag  is not closed at the same line")
		}

		tag := line[index+len(PREFIX) : indexEnd]
		line = line[:index] + template.ParseTag(tag) + line[indexEnd+len(SUFFIX):]
		template.Write(line)
	}

	return template
}

func (template *Template) ParseTag(tag string) string {
	var attributes map[string]string
	args := strings.SplitN(tag, " ", 2)
	command := args[0]

	if len(args) > 1 {
		attributes = parseAttributes(args[1])
	}

	switch command {
	case "region":
		template.writingTo = args[1]
		return ""
	case "insert":
		parts := strings.Split(strings.SplitN(args[1], " ", 2)[0], ":")
		if len(parts) == 2 {
			widget, ok := COMPONENTS[parts[0]]
			if !ok {
				fmt.Printf("WARNING: Couldn't find %q component.\n", parts[0])
				return ""
			}
			section, ok := widget[parts[1]]
			if !ok {
				fmt.Printf(
					"WARNING: Couldn't find section %q inside %q component.\n",
					parts[1],
					parts[0],
				)
				return ""
			}

			for key, value := range attributes {
				section = strings.ReplaceAll(section,
					PREFIX+"insert "+key+SUFFIX,
					strings.TrimSpace(template.getReplacement(key, value)),
				)
			}

			return section
		}
	}

	fmt.Println(command, args)
	return ""
}

func (template *Template) getReplacement(key string, value string) string {
	switch key {
	case "icon":
		icon, ok := ICONS[value]
		if !ok {
			fmt.Printf("WARNING: Unknown icon %q\n", value)
		}
		return icon
	default:
		return value
	}
}

func parseAttributes(input string) map[string]string {
	result := make(map[string]string)

	pairs := splitBySpacesOutsideQuotes(input)

	for _, pair := range pairs {
		attr := strings.SplitN(pair, "=", 2)
		if len(attr) == 2 {
			key := attr[0]
			value := strings.Trim(attr[1], "\"")
			result[key] = value
		}
	}

	return result
}

// Helper function to split by spaces that are outside of quotes
func splitBySpacesOutsideQuotes(s string) []string {
	var result []string
	var current strings.Builder
	inQuotes := false

	for _, r := range s {
		switch r {
		case ' ':
			if !inQuotes {
				if current.Len() > 0 {
					result = append(result, current.String())
					current.Reset()
				}
			} else {
				current.WriteRune(r)
			}
		case '"':
			inQuotes = !inQuotes
			current.WriteRune(r)
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}
