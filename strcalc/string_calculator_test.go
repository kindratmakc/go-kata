package strcalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumEmptyStringIsZero(t *testing.T) {
	assert.Equal(t, 0, Sum(""))
}

func TestSumOneNumber(t *testing.T) {
	assert.Equal(t, 1, Sum("1"))
}

func TestSumManyNumberDelimitedByComma(t *testing.T) {
	assert.Equal(t, 55, Sum("5,10,15,25"))
}

func TestSumNumbersDelimitedByCommaAndNewLine(t *testing.T) {
	assert.Equal(t, 7, Sum("1\n2,4"))
}
