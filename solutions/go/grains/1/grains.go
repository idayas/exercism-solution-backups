package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, errors.New("Number out of range")
	}
	return uint64(math.Pow(2, float64(number) - 1)), nil
}

func Total() uint64 {
	return uint64(math.Pow(2, 64))
}
