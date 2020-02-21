package strcalc

import (
	"strconv"
	"strings"
)

func Sum(input string) int {
	return sum(split(input))
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func split(input string) []int {
	var result []int
	byComma := strings.Split(input, ",")
	for _, number := range byComma {
		if hasNewLine(number) {
			byNewLine := strings.Split(number, "\n")
			for _, number := range byNewLine {
				result = addToSlice(number, result)

			}
		} else {
			result = addToSlice(number, result)
		}

	}

	return result
}

func addToSlice(number string, result []int) []int {
	integer, _ := strconv.Atoi(number)

	return append(result, integer)
}

func hasNewLine(number string) bool {
	return -1 != strings.Index(number, "\n")
}
