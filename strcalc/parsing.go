package strcalc

import (
	"strings"
)

func split(input string) []string {
	if strings.Count(input, "[") > 1 {
		delimiters := getDelimiters(input)
		withoutDelimiters := removeDelimiter(input)
		byDelimiters := splitByDelimiters(withoutDelimiters, delimiters)

		return byDelimiters
	}
	if hasOptionalDelimiter(input) {
		delimiter := getDelimiter(input)
		withoutDelimiter := removeDelimiter(input)
		byDelimiter := splitByDelimiter(withoutDelimiter, delimiter)

		return byDelimiter
	}

	return splitByDelimiters(input, []string{",", "\n"})
}

func splitByDelimiters(input string, delimiters []string) []string {
	var res []string
	buf := ""
	for _, c := range input {
		char := string(c)
		if isDelimiter(char, delimiters) {
			res = append(res, buf)
			buf = ""
			continue
		}
		buf += char
	}
	if len(buf) > 0 {
		res = append(res, buf)
	}

	return res
}

func isDelimiter(char string, delims []string) bool {
	for _, delim := range delims {
		if char == delim {
			return true
		}
	}

	return false
}

func splitByDelimiter(input, delimiter string) []string {
	var result []string
	byDelimiter := strings.Split(input, delimiter)
	for _, number := range byDelimiter {
		result = append(result, number)
	}

	return result
}

func hasOptionalDelimiter(input string) bool {
	return 0 == strings.Index(input, "//")
}

func removeDelimiter(input string) string {
	newLineAt := strings.Index(input, "\n")

	return input[newLineAt+1:]
}

func getDelimiter(input string) string {
	newLineAt := strings.Index(input, "\n")

	delimiter := input[2:newLineAt]
	firstChar := delimiter[0:1]
	lastChar := delimiter[len(delimiter)-1:]
	isMultiSymbolDelimiter := firstChar == "[" && lastChar == "]"

	if isMultiSymbolDelimiter {
		return delimiter[1 : len(delimiter)-1]
	}

	return delimiter
}

func getDelimiters(input string) []string {
	newLineAt := strings.Index(input, "\n")

	var res []string
	delimiters := strings.Split(input[2:newLineAt], "]")
	for _, elem := range delimiters {
		if elem == "" {
			break
		}
		res = append(res, elem[1:])
	}

	return res
}
