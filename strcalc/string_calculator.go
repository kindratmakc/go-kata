package strcalc

// https://kata-log.rocks/string-calculator-kata
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
