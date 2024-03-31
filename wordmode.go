package babble

import "strings"

// WordMode is the default means of processing key material. It splits
// on whitespace and uses WordTokens.
type WordMode struct {
	skip int
}

// Divider returns the divider to use between words.
func (m *WordMode) Divider() string {
	return " "
}

// Skip will cause Split() to skip the first n tokens.
func (m *WordMode) Skip(n int) {
	m.skip = n
}

// Tokenize will split on any whitespace.
func (m *WordMode) Tokenize(b []byte) []Token {
	var out []Token

	for _, word := range strings.Fields(string(b))[m.skip:] {
		out = append(out, WordToken{word})
	}

	return out
}
