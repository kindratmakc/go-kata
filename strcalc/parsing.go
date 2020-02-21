package strcalc

import "strings"

func split(input string) []string {
	if hasOptionalDelimiter(input) {
		delimiter := getDelimiter(input)
		withoutDelimiter := removeDelimiter(input)
		byDelimiter := splitByDelimiter(withoutDelimiter, delimiter)

		return byDelimiter
	}

	var result []string
	byComma := splitByDelimiter(input, ",")
	for _, n := range byComma {
		andByNewLine := splitByDelimiter(n, "\n")
		result = append(result, andByNewLine...)
	}

	return result
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
	delimiterAt := strings.Index(input, getDelimiter(input))

	return input[delimiterAt+2:]
}

func getDelimiter(input string) string {
	newLineAt := strings.Index(input, "\n")

	return input[2:newLineAt]
}
