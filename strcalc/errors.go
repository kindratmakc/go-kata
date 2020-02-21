package strcalc

import (
	"fmt"
	"strconv"
	"strings"
)

type NegativesNotAllowedError struct {
	Negatives []int
}

func (n NegativesNotAllowedError) Error() string {
	return fmt.Sprintf("negatives not allowed: %s", strings.Join(n.negativesAsStrings(), ","))
}

func (n NegativesNotAllowedError) negativesAsStrings() []string {
	var res []string
	for _, negative := range n.Negatives {
		res = append(res, strconv.Itoa(negative))
	}

	return res
}
