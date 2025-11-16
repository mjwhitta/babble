package babble

import (
	"fmt"
	"slices"
	"strings"
)

// StringToken normalizes by removing extra whitespace, converting to
// lowercase alphanumerical.
type StringToken struct {
	string

	symbols bool
}

// NewStringToken will return a StringToken.
func NewStringToken(s string, symbols ...bool) StringToken {
	if len(symbols) == 0 {
		symbols = []bool{false}
	}

	return StringToken{s, symbols[0]}.normalize()
}

// Bytes will return a []byte representation of the StringToken.
func (t StringToken) Bytes() []byte {
	return []byte(t.string)
}

// normalize will allow for obfuscation.
func (t StringToken) normalize() StringToken {
	var keep []rune = []rune{
		' ', '!', '"', '\'', ',', '-', '.', ':', ';', '?',
	}
	var s string = reWhiteSpace.ReplaceAllString(t.string, " ")

	s = strings.Map(
		func(r rune) rune {
			switch {
			case (r >= '0') && (r <= '9'):
				return r
			case (r >= 'a') && (r <= 'z'):
				return r
			}

			if t.symbols && slices.Contains(keep, r) {
				return r
			}

			return -1
		},
		strings.ToLower(strings.TrimSpace(s)),
	)

	// Some symbols may have been deleted, creating larger whitespace
	// blobs or leading/trailing whitespace
	s = strings.TrimSpace(reWhiteSpace.ReplaceAllString(s, " "))

	return StringToken{s, t.symbols}
}

// String will return a string representation of the StringToken.
func (t StringToken) String() string {
	if !t.symbols {
		return fmt.Sprintf("babble.NewStringToken(`%s`)", t.string)
	}

	return fmt.Sprintf("babble.NewStringToken(`%s`, true)", t.string)
}

// Valid is true for all non-empty StringTokens.
func (t StringToken) Valid() bool {
	return t.string != ""
}
