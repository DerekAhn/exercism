package twofer

import "fmt"

// ShareWith returns a conditional string: [name] || "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
