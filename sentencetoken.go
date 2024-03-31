package babble

import (
	"strings"

	hl "github.com/mjwhitta/hilighter"
)

// SentenceToken normalizes by removing extra whitespace and special
// characters.
type SentenceToken struct {
	string
}

// NewSentenceToken will return a SentenceToken.
func NewSentenceToken(s string) Token {
	return SentenceToken{s}
}

// Bytes will return a []byte representation of the SentenceToken.
func (t SentenceToken) Bytes() []byte {
	return []byte(t.string)
}

// Normalize will remove extra whitespace and special characters.
func (t SentenceToken) Normalize() Token {
	var s string = whiteSpace.ReplaceAllString(t.string, " ")

	s = strings.Map(
		func(r rune) rune {
			var keep []rune = []rune{
				' ', '!', '"', '\'', ',', '-', '.', ':', ';', '?',
			}

			if r == '`' {
				return '\''
			} else if (r >= '0') && (r <= '9') {
				return r
			} else if (r >= 'A') && (r <= 'Z') {
				return r
			} else if (r >= 'a') && (r <= 'z') {
				return r
			}

			for i := range keep {
				if r == keep[i] {
					return r
				}
			}

			return -1
		},
		strings.TrimSpace(s),
	)

	if s == "" {
		return nil
	}

	return SentenceToken{s}
}

// String will return a string representation of the SentenceToken.
func (t SentenceToken) String() string {
	return hl.Sprintf("babble.NewSentenceToken(`%s`)", t.string)
}
