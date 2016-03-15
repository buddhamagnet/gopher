package numbers

import (
	"fmt"
	"math"
)

// EqualFloat compares two floats for a given
// accuracy - or to the greatest accuracy the
// machine can achieve if a negative limit is
// passed in.
func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <=
		(limit * math.Min(math.Abs(x), math.Abs(y)))
}

// EqualsFloatPrec checks for float equality using
// string conversions - a slower approach than EqualFloat.
func EqualFloatPrec(x, y float64, decimals int) bool {
	a := fmt.Sprintf("%.*f", decimals, x)
	b := fmt.Sprintf("%.*f", decimals, y)
	return len(a) == len(b) && a == b
}

// IntFromFloat64 safely downsizes floats to integers.
func IntFromFloat64(x float64) (int, error) {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole), nil
	}
	return fmt.Errorf("%g is out of int32 range", x)
}

// Uint8FromInt safely downsizes integers to bytes.
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of uint8 range", x)
}
