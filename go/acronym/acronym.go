package acronym

import "strings"

// Abbreviate converts a phrase to its arconym
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Trim(s, " ")

	var a string
	for _, ss := range strings.Split(s, " ") {
		a += strings.ToUpper(string(ss[0]))
	}
	return a
}
