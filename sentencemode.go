package babble

import (
	"regexp"
	"strings"
)

// SentenceMode will process key material by splitting on typical
// sentence-ending punctuation. It uses SentenceTokens.
type SentenceMode struct {
	skip int
}

// Divider returns the divider to use between sentences.
func (m *SentenceMode) Divider() string {
	return " "
}

// Skip will cause Split() to skip the first n tokens.
func (m *SentenceMode) Skip(n int) {
	m.skip = n
}

// Tokenize will split on any typical sentence-ending punctuation.
func (m *SentenceMode) Tokenize(b []byte) []Token {
	var out []Token
	var r *regexp.Regexp = regexp.MustCompile(
		`([A-Z][^!"'.?]*(("[^"]+"|'[^']+')?[^!"'.?]*)*([!.?])+)`,
	)
	var s string = whiteSpace.ReplaceAllString(string(b), " ")

	s = strings.TrimSpace(s)

	for _, sentence := range r.FindAllString(s, -1)[m.skip:] {
		out = append(out, SentenceToken{sentence})
	}

	return out
}
