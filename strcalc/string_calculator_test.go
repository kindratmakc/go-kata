package strcalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumEmptyStringIsZero(t *testing.T) {
	assert.Equal(t, 0, Sum(""))
}
