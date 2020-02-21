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

func TestSumTwoNumbersDelimitedByComma(t *testing.T) {
	assert.Equal(t, 3, Sum("1,2"))
}
