package strcalc

import (
	"strconv"
	"strings"
)

func Sum(input string) (int, error) {
	numbers := split(input)
	integers := toInt(numbers)
	err := checkForNegatives(integers)
	if nil != err {
		return 0, err
	}

	return sum(integers), nil
}

func checkForNegatives(integers []int) error {
	var negatives []int
	for _, integer := range integers {
		if integer < 0 {
			negatives = append(negatives, integer)
		}
	}

	if len(negatives) > 0 {
		return NegativesNotAllowedError{Negatives: negatives}
	}

	return nil
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
