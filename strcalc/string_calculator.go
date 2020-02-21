package strcalc

import (
	"strconv"
	"strings"
)

func Sum(input string) (int, error) {
	numbers := split(input)
	ints := toInt(numbers)
	ints = removeGraterThan1000(ints)
	err := checkForNegatives(ints)
	if nil != err {
		return 0, err
	}

	return sum(ints), nil
}

func removeGraterThan1000(ints []int) []int {
	return filterInts(ints, func(i int) bool {
		return i <= 1000
	})
}

func checkForNegatives(integers []int) error {
	negatives := filterInts(integers, func(i int) bool {
		return i < 0
	})

	if len(negatives) > 0 {
		return NegativesNotAllowedError{Negatives: negatives}
	}

	return nil
}

func filterInts(ints []int, filter func(int) bool) []int {
	var filtered []int
	for _, i := range ints {
		if filter(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func toInt(numbers []string) []int {
	var result []int
	for _, number := range numbers {
		integer, _ := strconv.Atoi(number)
		result = append(result, integer)
	}

	return result
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

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
