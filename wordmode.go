package babble

import "strings"

// WordMode is the default means of processing key material. It splits
// on whitespace and uses StringTokens.
type WordMode struct {
	offset uint
}

// AllowsMultiples is true for WordMode.
func (m *WordMode) AllowsMultiples() bool {
	return true
}

// Divider returns the divider to use between words.
func (m *WordMode) Divider() string {
	return " "
}

// Seek will cause Split() to seek to the specified word.
func (m *WordMode) Seek(n uint) {
	m.offset = n
}

// Tokenize will split on any whitespace.
func (m *WordMode) Tokenize(b []byte) []Token {
	var offset uint = m.offset
	var out []Token

	if offset > uint(len(b)) {
		offset = 0
	}

	for _, word := range strings.Fields(string(b))[offset:] {
		out = append(out, NewStringToken(word))
	}

	return out
}
