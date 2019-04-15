// Package luhn validates strings per Luhn alogrithm
package luhn

import "regexp"
import "github.com/golang/example/stringutil"

// method Valid accepts a string as input and returns a boolean for validity check
func Valid(input string) bool {
	if matched, _ := regexp.Match(`[^0-9 ]+`, []byte(input)); matched {
		return false
	}
	reg, _ := regexp.Compile("[ ]+")
	cleaned := reg.ReplaceAllString(input, "")
	if len(cleaned) < 2 {
		return false
	}

	var reversed = stringutil.Reverse(cleaned)

	checksum := 0
	for i, c := range reversed {
		cAsInt := int(c - '0')
		if i%2 != 0 {
			doubled := cAsInt * 2
			if doubled > 9 {
				checksum += (doubled - 9)
			} else {
				checksum += doubled
			}
		} else {
			checksum += cAsInt
		}
	}
	return checksum%10 == 0
}
