package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var valid, output, err = validate(inputBase, inputDigits, outputBase)
	if !valid {
		return output, err
	}
	var decimalTotal = baseToDecimal(inputBase, inputDigits)
	if decimalTotal == 0 {
		return []int{0}, nil
	}
	return decimalToBase(outputBase, decimalTotal), nil
}

func validate(inputBase int, inputDigits []int, outputBase int) (bool, []int, error) {
	if inputBase < 2 {
		return false, []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return false, []int{}, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		return false, []int{0}, nil
	}
	for _, digit := range inputDigits {
		if 0 > digit || digit >= inputBase {
			return false, []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	return true, nil, nil
}

func baseToDecimal(base int, digits []int) int {
	var decimalTotal = 0
	var maxPower = len(digits) - 1
	for place, digit := range digits {
		var converted = float64(digit) * math.Pow(float64(base), float64(maxPower-place))
		decimalTotal += int(converted)
	}
	return decimalTotal
}

func decimalToBase(base int, value int) []int {
	var output = []int{}
	var remainder = value
	for {
		if remainder > 0 {
			output = append(output, remainder%base)
			remainder = int(math.Floor(float64(remainder / base)))
		} else {
			break
		}
	}
	return reverse(output)
}

func reverse(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var maxIndex = len(arr) - 1
	for idx := 0; idx <= maxIndex/2; idx++ {
		var tmp = arr[idx]
		arr[idx] = arr[maxIndex-idx]
		arr[maxIndex-idx] = tmp
	}
	return arr
}
