package babble

// Mode is an interface that allows for customization of how a Key is
// created from key material and how ciphertext/plaintext is
// processed.
type Mode interface {
	Divider() string
	Skip(int)
	Tokenize([]byte) []Token
}
