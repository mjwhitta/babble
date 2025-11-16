package babble

import (
	"regexp"
)

// ParagraphMode will process key material by splitting on typical
// paragraph breaks (two or more newlines). It uses StringTokens.
type ParagraphMode struct {
	offset int
}

// AllowsMultiples is true for ParagraphMode.
func (m *ParagraphMode) AllowsMultiples() bool {
	return true
}

// Divider returns the divider to use between paragraphs.
func (m *ParagraphMode) Divider() string {
	return "\n\n"
}

// Seek will cause Split() to seek to the specified paragraph.
func (m *ParagraphMode) Seek(n int) {
	m.offset = n
}

// Tokenize will split on a typical paragraph break (two or more
// newlines).
func (m *ParagraphMode) Tokenize(b []byte) []Token {
	var offset int = m.offset
	var out []Token
	var r *regexp.Regexp = regexp.MustCompile(`\n\n+`)

	if offset > len(b) {
		offset = 0
	}

	for _, paragraph := range r.Split(string(b), -1)[offset:] {
		out = append(out, NewStringToken(paragraph, true))
	}

	return out
}
