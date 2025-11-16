package babble

// ByteMode is the default means of processing key material. It splits
// on whitespace and uses ByteTokens.
type ByteMode struct {
	offset int
}

// AllowsMultiples is false for ByteMode.
func (m *ByteMode) AllowsMultiples() bool {
	return false
}

// Divider returns the divider to use between bytes.
func (m *ByteMode) Divider() string {
	return ""
}

// Seek will cause Split() to seek to the specified byte.
func (m *ByteMode) Seek(n int) {
	m.offset = n
}

// Tokenize will split on any whitespace.
func (m *ByteMode) Tokenize(b []byte) []Token {
	var offset int = m.offset
	var out []Token

	if offset > len(b) {
		offset = 0
	}

	for _, char := range b[offset:] {
		out = append(out, NewByteToken(char))
	}

	return out
}
