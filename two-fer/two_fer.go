// Package twofer shares stuff
package twofer

import "fmt"

// ShareWith takes an optional name.  If name is not empty, shares between 'me' and 'name'
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
