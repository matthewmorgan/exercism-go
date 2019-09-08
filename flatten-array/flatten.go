package flatten

import "reflect"

// Method Flatten takes an array of arbitrary types, which could be mixed or arrays themselves,
// and returns a flat array.
func Flatten(nested interface{}) []interface{} {
	// Make an array of interfaces to receive the flattened elements.
	flattened := make([]interface{}, 0)
	// Check if our input contains a nested slice.
	nestedSlice , ok := nested.([]interface{})
	// If not, return the input as is (it's already flat!)
	if !ok {
		return flattened
	}
	// If nested, then iterate over the elements of the nested slice
	for _, e := range nestedSlice {
		// Use reflection API to introspect elements.
		t := reflect.TypeOf(e)
		// Throw out nils.
		if t == nil {
			continue
		}
		switch kind := t.Kind(); kind {
		// If our element is itself a slice, call Flatten on it and append
		// the results.
		case reflect.Slice:
			flattened = append(flattened, Flatten(e)...)
		// Otherwise just append the already flat element.
		default:
			flattened = append(flattened, e)
		}
	}
	// Return the final result.
	return flattened
}
