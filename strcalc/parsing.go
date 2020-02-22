package strcalc

import (
	"strings"
)

func split(input string) []string {
	delimiters, body := parse(input)

	return splitByDelimiters(delimiters, body)
}

func parse(input string) ([]string, string) {
	if hasDelimitersDeclaration(input) {
		delimiters, body := separate(input)

		return parseDelimiters(delimiters), body
	}
	return defaultDelimiters(), input
}

func splitByDelimiters(delimiters []string, input string) []string {
	buf := input
	for _, delimiter := range delimiters {
		buf = strings.ReplaceAll(buf, delimiter, ",")
	}

	return strings.Split(buf, ",")
}

func hasDelimitersDeclaration(input string) bool {
	return strings.Contains(input, "//")
}

func separate(input string) (string, string) {
	newLineAt := strings.Index(input, "\n")
	delimiters := input[2:newLineAt]
	body := input[newLineAt+1:]

	return delimiters, body
}

func parseDelimiters(input string) []string {
	return strings.FieldsFunc(input, func(c rune) bool {
		return c == ']' || c == '['
	})
}

func defaultDelimiters() []string {
	return []string{",", "\n"}
}
