package babble

// ByteMode is the default means of processing key material. It splits
// on whitespace and uses ByteTokens.
type ByteMode struct {
	skip int
}

// Divider returns the divider to use between bytes.
func (m *ByteMode) Divider() string {
	return ""
}

// Skip will cause Split() to skip the first n tokens.
func (m *ByteMode) Skip(n int) {
	m.skip = n
}

// Tokenize will split on any whitespace.
func (m *ByteMode) Tokenize(b []byte) []Token {
	var out []Token

	for _, char := range b[m.skip:] {
		out = append(out, ByteToken{char})
	}

	return out
}
