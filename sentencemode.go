package babble

import "regexp"

// SentenceMode will process key material by splitting on typical
// sentence-ending punctuation. It uses StringTokens.
type SentenceMode struct {
	offset uint
}

// AllowsMultiples is true for SentenceMode.
func (m *SentenceMode) AllowsMultiples() bool {
	return true
}

// Divider returns the divider to use between sentences.
func (m *SentenceMode) Divider() string {
	return " "
}

// Seek will cause Split() to seek to the specified sentence.
func (m *SentenceMode) Seek(n uint) {
	m.offset = n
}

// Tokenize will split on any typical sentence-ending punctuation.
func (m *SentenceMode) Tokenize(b []byte) []Token {
	var offset uint = m.offset
	var out []Token
	var r *regexp.Regexp = regexp.MustCompile(
		`([A-Za-z][^!"'.?]+[!.?]|"[^"]+[!.?]"|'.*?[!.?]')`,
	)
	var s string = reWhiteSpace.ReplaceAllString(string(b), " ")

	if offset > uint(len(b)) {
		offset = 0
	}

	for _, sentence := range r.FindAllString(s, -1)[offset:] {
		out = append(out, NewStringToken(sentence, true))
	}

	return out
}
