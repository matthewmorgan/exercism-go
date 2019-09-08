package accumulate

// Method Accumulate accepts an input collection and an operation to perform on each
// element in the collection, returning a collection as a result.
func Accumulate(input []string, operation func(string) string) []string {
	output := make([]string, len(input))
	for idx, el := range input {
		output[idx] = operation(el)
	}
	return output
}