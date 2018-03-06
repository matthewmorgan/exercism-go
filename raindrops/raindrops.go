// Package raindrops contains a function to convert numbers into raindrop-speak
package raindrops

import (
	"sort"
	"strconv"
)

var RAINDROP_RESPONSES = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

var KEYS []int

func init() {
	for k := range RAINDROP_RESPONSES {
		KEYS = append(KEYS, k)
	}
	sort.Ints(KEYS)
}

// Convert takes an integer, and returns a response according to the factors of the input
func Convert(input int) string {
	response := ""
	for i := 0; i < len(KEYS); i++ {
		key := KEYS[i]
		if input%key == 0 {
			response += RAINDROP_RESPONSES[key]
		}
	}
	if len(response) == 0 {
		return strconv.Itoa(input)
	}
	return response
}
