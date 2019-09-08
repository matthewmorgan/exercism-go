// Package triangle determines the type of a triangle based on the lengths of its sides
package triangle

import (
	"errors"
	"math"
)

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
)

// KindFromSides returns the Kind of a triangle based on the lengths of its sides
func KindFromSides(a, b, c float64) Kind {
	if paramsIllegal(a, b, c) != nil {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

// Validates parameters, returning error if any are illegal
func paramsIllegal(a, b, c float64) error {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return errors.New("NaN not allowed")
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return errors.New("infinite values not allowed")
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return errors.New("negative values not allowed")
	}
	if a+b < c || a+c < b || b+c < a {
		return errors.New("parameters violate triangle inequality")
	}
	return nil
}