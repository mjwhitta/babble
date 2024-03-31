package babble

import (
	"regexp"
	"strings"
)

// ParagraphMode will process key material by splitting on typical
// paragraph breaks (two or more newlines). It uses ParagraphTokens.
type ParagraphMode struct {
	skip int
}

// Divider returns the divider to use between paragraphs.
func (m *ParagraphMode) Divider() string {
	return "\n\n"
}

// Skip will cause Split() to skip the first n paragraphs.
func (m *ParagraphMode) Skip(n int) {
	m.skip = n
}

// Tokenize will split on a typical paragraph break (two or more
// newlines), skipping the first n paragraphs.
func (m *ParagraphMode) Tokenize(b []byte) []Token {
	var out []Token
	var r *regexp.Regexp = regexp.MustCompile(`\n\n+`)

	for _, paragraph := range strings.Split(
		r.ReplaceAllString(string(b), "\n\n"), "\n\n",
	)[m.skip:] {
		out = append(out, ParagraphToken{paragraph})
	}

	return out
}
