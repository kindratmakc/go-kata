package strcalc

import (
	"strconv"
	"strings"
)

func Sum(numbers string) int {
	result := 0
	splitted := strings.Split(numbers, ",")
	for _, number := range splitted {
		newLineIndex := strings.Index(number, "\n")
		if -1 == newLineIndex {
			integer, _ := strconv.Atoi(number)
			result += integer
		} else {
			subSplit := strings.Split(number, "\n")
			for _, subNumber := range subSplit {
				integer, _ := strconv.Atoi(subNumber)
				result += integer

			}
		}

	}

	return result
}
