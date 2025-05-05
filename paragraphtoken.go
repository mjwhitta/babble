package babble

import (
	"slices"
	"strings"

	hl "github.com/mjwhitta/hilighter"
)

// ParagraphToken normalizes by removing extra whitespace and special
// characters.
type ParagraphToken struct {
	string
}

// NewParagraphToken will return a ParagraphToken.
func NewParagraphToken(p string) Token {
	return ParagraphToken{p}
}

// Bytes will return a []byte representation of the ParagraphToken.
func (t ParagraphToken) Bytes() []byte {
	return []byte(t.string)
}

// Normalize will remove extra whitespace and special characters.
func (t ParagraphToken) Normalize() Token {
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

			if slices.Contains(keep, r) {
				return r
			}

			return -1
		},
		strings.TrimSpace(s),
	)

	if s == "" {
		return nil
	}

	return ParagraphToken{s}
}

// String will return a string representation of the ParagraphToken.
func (t ParagraphToken) String() string {
	return hl.Sprintf("babble.NewParagraphToken(`%s`)", t.string)
}
