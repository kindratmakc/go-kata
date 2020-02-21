package strcalc

import "strconv"

func Sum(numbers string) int {
	result, _ := strconv.Atoi(numbers)

	return result
}
