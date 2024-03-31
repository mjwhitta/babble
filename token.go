package babble

// Token is an interface that allows for processing of key material
// and ciphertext/plaintext. It has the ability to normalize itself
// allowing for obfuscation of the ciphertext.
type Token interface {
	Bytes() []byte
	Normalize() Token
	String() string
}
