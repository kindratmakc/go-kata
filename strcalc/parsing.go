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
