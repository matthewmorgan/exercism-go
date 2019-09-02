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

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func ConcurrentFrequency(languageStrings []string) FreqMap {
	mapsChannel := make(chan FreqMap)
	for _, languageString := range languageStrings {
		// Using a lambda function here.  Not sure I like this better than using a named function.
		go func(s string, c chan FreqMap) {
			c <- Frequency(s)
		}(languageString, mapsChannel)
	}

	allLanguageCounts := FreqMap{}
	for range languageStrings {
		languageMap := <- mapsChannel
		for r, count := range languageMap {
			allLanguageCounts[r] += count
		}
	}
	return allLanguageCounts
}
