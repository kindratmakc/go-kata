package strcalc

import "strconv"

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
	var res []int
	for _, i := range ints {
		if filter(i) {
			res = append(res, i)
		}
	}
	return res
}

func toInt(numbers []string) []int {
	var res []int
	for _, number := range numbers {
		i, _ := strconv.Atoi(number)
		res = append(res, i)
	}

	return res
}

func sum(ints []int) int {
	res := 0
	for _, i := range ints {
		res += i
	}

	return res
}
