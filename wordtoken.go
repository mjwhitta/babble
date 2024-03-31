package babble

import (
	"strings"

	hl "github.com/mjwhitta/hilighter"
)

// WordToken normalizes to lowercase alphanumerical.
type WordToken struct {
	string
}

// NewWordToken will return a WordToken.
func NewWordToken(w string) Token {
	return WordToken{w}
}

// Bytes will return a []byte representation of the WordToken.
func (t WordToken) Bytes() []byte {
	return []byte(t.string)
}

// Normalize will remove non-alphanumeric characters and convert to
// lowercase.
func (t WordToken) Normalize() Token {
	var s string = strings.Map(
		func(r rune) rune {
			if (r >= '0') && (r <= '9') {
				return r
			} else if (r >= 'a') && (r <= 'z') {
				return r
			}

			return -1
		},
		strings.ToLower(t.string),
	)

	if s == "" {
		return nil
	}

	return WordToken{s}
}

// String will return a string representation of the WordToken.
func (t WordToken) String() string {
	return hl.Sprintf("babble.NewWordToken(`%s`)", t.string)
}
