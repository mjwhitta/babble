package babble

import (
	"os"
	"strings"

	"github.com/mjwhitta/errors"
	hl "github.com/mjwhitta/hilighter"
)

// Key is a struct containing the key data required to decrypt/encrypt
// payloads.
type Key struct {
	Mode Mode

	key    map[byte][]Token
	revkey map[Token]byte
	width  int
}

// newKey will return a pointer to a new Key.
func newKey(width ...int) *Key {
	if (len(width) == 0) || width[0] < 1 {
		width = []int{1}
	}

	return &Key{
		key:    make(map[byte][]Token, 256),
		Mode:   &WordMode{},
		revkey: make(map[Token]byte, width[0]*256),
		width:  width[0],
	}
}

// NewKeyFromBytes will parse a byte array and return a pointer to a
// new Key.
func NewKeyFromBytes(b []byte, m Mode, width ...int) (*Key, error) {
	var e error
	var i int
	var k *Key = newKey(width...)

	k.Mode = m

	for _, t := range k.Mode.Tokenize(b) {
		// Allow for obfuscation of key file
		if t = t.Normalize(); t == nil {
			continue
		}

		if e = k.Set(byte(i%256), t); e != nil {
			// Allow for obfuscation of key file
			continue
		}

		i++

		// Ignore extras
		if i == (k.width * 256) {
			break
		}
	}

	if i != (k.width * 256) {
		e = errors.Newf("key file missing %d tokens", (k.width*256)-i)
		return nil, e
	}

	return k, nil
}

// NewKeyFromFile will read in a file and return a pointer to a new
// Key.
func NewKeyFromFile(fn string, m Mode, width ...int) (*Key, error) {
	var b []byte
	var e error

	if b, e = os.ReadFile(fn); e != nil {
		return nil, errors.Newf("failed to read %s: %t", fn, e)
	}

	return NewKeyFromBytes(b, m, width...)
}

// NewKeyFromMap will accept an already created map and will return a
// pointer to a new Key instance.
func NewKeyFromMap(mapping map[byte][]Token, m Mode) (*Key, error) {
	var k *Key = newKey()

	k.key = mapping
	k.Mode = m
	k.revkey = map[Token]byte{}

	if len(mapping) != 256 {
		return nil, errors.New("key is missing entries")
	}

	for b, s := range mapping {
		if width := len(s); width > 0 {
			k.width = width
		} else {
			return nil, errors.Newf("key is missing entry: %02x", b)
		}

		for i := range s {
			k.revkey[s[i]] = b
		}
	}

	return k, nil
}

// ByteFor will return the byte associated with the specified token.
func (k *Key) ByteFor(t Token) (b byte, ok bool) {
	// Allow for obfuscation of token
	b, ok = k.revkey[t.Normalize()]
	return
}

// Set will link the specified byte and token for use with
// decryption/encryption.
func (k *Key) Set(b byte, t Token) error {
	// Allow for obfuscation of token
	t = t.Normalize()

	if _, ok := k.revkey[t]; ok {
		return errors.Newf("key includes duplicate token: %s", t)
	}

	k.key[b] = append(k.key[b], t)
	k.revkey[t] = b

	return nil
}

// String will return a string representation of the Key.
func (k *Key) String() string {
	var out []string = []string{
		"package main\n",
		"import \"github.com/mjwhitta/babble\"\n",
		"func babbleDecrypt(b []byte) ([]byte, error) {",
		"\tvar k *babble.Key",
		"",
		"\tk, _ = babble.NewKeyFromMap(",
		"\t\tmap[byte][]babble.Token{",
	}

	for i := range 256 {
		out = append(out, hl.Sprintf("\t\t\t0x%02x: {", i))

		for _, t := range k.key[byte(i)] {
			out = append(out, hl.Sprintf("\t\t\t\t%s,", t.String()))
		}

		out = append(out, "\t\t\t},")
	}

	out = append(
		out,
		"\t\t},",
		strings.ReplaceAll(hl.Sprintf("\t\t&%T{},", k.Mode), "*", ""),
		"\t)",
		"",
		"\treturn babble.Decrypt(b, k)",
		"}",
	)

	return strings.Join(out, "\n")
}

// TokenFor will return the token associated with the specified byte.
func (k *Key) TokenFor(b byte) (Token, error) {
	if len(k.key[b]) == 0 {
		return nil, errors.Newf("key missing entry for: %02x", b)
	}

	if len(k.key[b]) == 1 {
		return k.key[b][0], nil
	}

	// Randomly pick one
	return k.key[b][randIntn(len(k.key[b]))], nil
}
