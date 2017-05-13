// Package cmdln will parse a command line string into two parts, using the bash argv parsing rules. It is intended to be passed to the exec.Command function.
package cmdln

import (
	"fmt"
	"strings"
)

const item = `'"\`

// line keeps track of spaces inside and outside of quotes
// it also track escaped quotes, so they won't effect word splitting rules
type line struct{ q, qq, bs bool }

// SplitFunc can be used to split strings on spaces that respect spaces within single and double quotes.
func (ln *line) SplitFunc(c rune) bool {
	switch {
	case c == rune(item[0]) && !ln.bs && !ln.qq:
		ln.q = !ln.q
	case c == rune(item[1]) && !ln.bs && !ln.q:
		ln.qq = !ln.qq
	case c == rune(item[2]) && !ln.bs:
		ln.bs = true
		return false
	case c == ' ' && !ln.q && !ln.qq:
		return true
	}

	ln.bs = false
	return false
}

// Split takes a string and splits it into two parts; command and a slice of arguments using bash word splitting rules
func Split(s string) (cmd string, args []string) {
	ln := &line{}
	ss := strings.FieldsFunc(s, ln.SplitFunc)
	return ss[0], ss[1:]
}

// Splitf formats according to a format specifier and splits the resulting string into two parts; a command and a slice of arguments using bash word splitting rules
func Splitf(format string, infs ...interface{}) (cmd string, args []string) {
	return Split(fmt.Sprintf(format, infs...))
}
