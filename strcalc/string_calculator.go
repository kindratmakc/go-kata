package strcalc

import (
	"strconv"
	"strings"
)

func Sum(numbers string) int {
	result := 0
	splitted := strings.Split(numbers, ",")
	for _, number := range splitted {
		integer, _ := strconv.Atoi(number)
		result += integer
	}

	return result
}
