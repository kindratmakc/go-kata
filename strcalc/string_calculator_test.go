package strcalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumEmptyStringIsZero(t *testing.T) {
	result, _ := Sum("")
	assert.Equal(t, 0, result)
}

func TestSumOneNumber(t *testing.T) {
	result, _ := Sum("1")
	assert.Equal(t, 1, result)
}

func TestSumManyNumberDelimitedByComma(t *testing.T) {
	result, _ := Sum("5,10,15,25")
	assert.Equal(t, 55, result)
}

func TestSumNumbersDelimitedByCommaAndNewLine(t *testing.T) {
	result, _ := Sum("1\n2,4")
	assert.Equal(t, 7, result)
}

func TestSumNumbersWithOptionalDelimiter(t *testing.T) {
	result, _ := Sum("//;\n1;2")
	assert.Equal(t, 3, result)
}

func TestSumReturnsErrorWithNegativeNumbers(t *testing.T) {
	_, err := Sum("-1,2,-2")
	assert.EqualError(t, err, "negatives not allowed: -1,-2")
}

func TestSumIgnoresNumbersGreaterThan1000(t *testing.T) {
	result, _ := Sum("5,1001,1000")
	assert.Equal(t, 1005, result)
}

func TestSumWithMultiSymbolDelimiter(t *testing.T) {
	result, _ := Sum("//[***]\n1***2***3")
	assert.Equal(t, 6, result)
}
