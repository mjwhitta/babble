package babble

import "fmt"

// ByteToken represetns a byte. There is no normalization.
type ByteToken struct {
	byte
}

// NewByteToken will return a ByteToken.
func NewByteToken(b byte) ByteToken {
	return ByteToken{b}
}

// Bytes will return a []byte representation of the ByteToken.
func (t ByteToken) Bytes() []byte {
	return []byte{t.byte}
}

// String will return a string representation of the ByteToken.
func (t ByteToken) String() string {
	return fmt.Sprintf("babble.NewByteToken(0x%02x)", t.byte)
}

// Valid is true for all ByteTokens.
func (t ByteToken) Valid() bool {
	return true
}
