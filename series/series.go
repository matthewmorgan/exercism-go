package series

func All(n int, s string) []string {
	results := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		results = append(results, slice(n, s[i:]))
	}
	return results
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool){
	if len(s) < n {
		return s, false
	}
	return UnsafeFirst(n, s), true
}

func slice(n int, s string) string {
	return s[0:n]
}
