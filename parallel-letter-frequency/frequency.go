package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates the letter frequency for an array of strings.
// Each string in the array represents a different language.
func ConcurrentFrequency(strings []string) FreqMap {
	// Create a map to aggregate counts by letter across all languages.
	sum := FreqMap{}
	// count is number of strings AKA languages.
	count := len(strings)
	// Create a channel to queue computed results for a given language.
	results := make(chan FreqMap, count)

	// Start a goroutine to calculate the frequency map for each lanugage,
	// sending the results to our channel AKA queue.
	for _, s := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	// Aggregate the frequency maps passed back via the results channel,
	// summing counts by letter across all languages.
	for i := 0; i < count; i++ {
		for r, freq := range <-results {
			sum[r] += freq
		}
	}

	return sum
}
