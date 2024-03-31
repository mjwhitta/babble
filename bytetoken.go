package babble

import hl "github.com/mjwhitta/hilighter"

// ByteToken represetns a byte. There is no normalization.
type ByteToken struct {
	byte
}

// NewByteToken will return a ByteToken.
func NewByteToken(b byte) Token {
	return ByteToken{b}
}

// Bytes will return a []byte representation of the ByteToken.
func (t ByteToken) Bytes() []byte {
	return []byte{t.byte}
}

// Normalize will remove non-alphanumeric characters and convert to
// lowercase.
func (t ByteToken) Normalize() Token {
	return t
}

// String will return a string representation of the ByteToken.
func (t ByteToken) String() string {
	return hl.Sprintf("babble.NewByteToken(0x%02x)", t.byte)
}
