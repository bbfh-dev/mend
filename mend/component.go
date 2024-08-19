package mend

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Component struct {
	inSection      bool
	currentSection string
	output         bytes.Buffer
	sections       map[string]*bytes.Buffer
}

func (component *Component) Output() map[string]string {
	var sections = map[string]string{}
	for key, value := range component.sections {
		sections[key] = value.String()
	}
	return sections
}

func ParseComponent(reader io.Reader) *Component {
	component := &Component{
		inSection:      false,
		currentSection: "",
		output:         bytes.Buffer{},
		sections:       map[string]*bytes.Buffer{},
	}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "<!-- begin ") {
			component.inSection = true
			component.currentSection = extractSectionName(line, "begin")
			component.sections[component.currentSection] = &bytes.Buffer{}
			continue
		}

		if strings.Contains(line, "<!-- end ") && component.inSection {
			sectionName := extractSectionName(line, "end")
			if sectionName == component.currentSection {
				component.inSection = false
				component.currentSection = ""
			}
			continue
		}

		if component.inSection {
			component.sections[component.currentSection].WriteString(line + "\n")
		} else if strings.Contains(line, "<!--") && strings.Contains(line, "-->") {
			// Replace other comments with a placeholder (or any other processing logic)
			component.output.WriteString(strings.ReplaceAll(strings.ReplaceAll(line, "<!--", ""), "-->", "") + "\n")
		} else {
			component.output.WriteString(line + "\n")
		}
	}

	return component
}

// extractSectionName extracts the section name from the comment
func extractSectionName(line, tag string) string {
	return strings.TrimSpace(
		strings.TrimSuffix(strings.TrimPrefix(line, fmt.Sprintf("<!-- %s ", tag)), "-->"),
	)
}
