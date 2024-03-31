package babble

import (
	"bytes"
	"os"

	"github.com/mjwhitta/errors"
)

// Decrypt will decrypt the provided []string using the provided Key.
func Decrypt(b []byte, k *Key) ([]byte, error) {
	var ptxt []byte

	for _, t := range k.Mode.Tokenize(b) {
		// Allow for obfuscation of the payload
		if tmp, ok := k.ByteFor(t); ok {
			ptxt = append(ptxt, tmp)
		}
	}

	if i := bytes.Index(ptxt, []byte(header[1:])); i >= 0 {
		ptxt = ptxt[i+len(header)-1:]
	} else {
		return nil, errors.New("failed to find babble header")
	}

	if i := bytes.Index(ptxt, []byte(footer)); i >= 0 {
		ptxt = ptxt[:i]
	} else {
		return nil, errors.New("failed to find babble footer")
	}

	return ptxt, nil
}

// DecryptFile will open the provided filename, read the contents, and
// decrypt using the provided Key.
func DecryptFile(fn string, k *Key) ([]byte, error) {
	var b []byte
	var e error

	if b, e = os.ReadFile(fn); e != nil {
		return nil, errors.Newf("decrypt failed: %w", e)
	}

	return Decrypt(b, k)
}

// Encrypt will encrypt the provided []byte using the provided Key.
func Encrypt(b []byte, k *Key) ([]byte, error) {
	var ctxt [][]byte

	b = append([]byte(header), b...)
	b = append(b, []byte(footer)...)

	for _, v := range b {
		if t, e := k.TokenFor(v); e != nil {
			return nil, e
		} else if t != nil {
			ctxt = append(ctxt, t.Bytes())
		}
	}

	return bytes.Join(ctxt, []byte(k.Mode.Divider())), nil
}

// EncryptFile will open the provided filename, read the contents, and
// encrypt using the provided Key.
func EncryptFile(fn string, k *Key) ([]byte, error) {
	var b []byte
	var e error

	if b, e = os.ReadFile(fn); e != nil {
		return nil, errors.Newf("encrypt failed: %w", e)
	}

	return Encrypt(b, k)
}
