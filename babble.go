package babble

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mjwhitta/errors"
	"github.com/mjwhitta/pathname"
)

// Decrypt will decrypt the provided []string using the provided Key.
func Decrypt(ctxt []byte, k *Key) ([]byte, error) {
	var ptxt []byte

	for _, t := range k.Mode.Tokenize(ctxt) {
		// Allow for obfuscation of the payload
		if tmp, ok := k.ByteFor(t); ok {
			ptxt = append(ptxt, tmp)
		}
	}

	if i := bytes.Index(ptxt, []byte(header)); i >= 0 {
		ptxt = ptxt[i+len(header):]
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

// DecryptCompare will decrypt the provided []string using the
// provided Key. It will then compare to the provided cmp. This is
// only intended for debugging.
func DecryptCompare(ctxt []byte, k *Key, cmp []byte) ([]byte, error) {
	var e error
	var j int
	var ptxt []byte

	if ptxt, e = Decrypt(ctxt, k); e != nil {
		return nil, e
	}

	for i := range ptxt {
		if j > len(cmp) {
			break
		}

		for ptxt[i] != cmp[j] {
			fmt.Printf("[%d] %02x != %02x\n", i, cmp[j], ptxt[i])
			j++
		}

		j++
	}

	return ptxt, nil
}

// DecryptFile will open the provided filename, read the contents, and
// decrypt using the provided Key.
func DecryptFile(fn string, k *Key) ([]byte, error) {
	var b []byte
	var e error

	if b, e = os.ReadFile(pathname.ExpandPath(fn)); e != nil {
		return nil, errors.Newf("failed to decrypt: %w", e)
	}

	return Decrypt(b, k)
}

// Encrypt will encrypt the provided []byte using the provided Key.
func Encrypt(ptxt []byte, k *Key) ([]byte, error) {
	var ctxt [][]byte

	ptxt = append([]byte(header), ptxt...)
	ptxt = append(ptxt, []byte(footer)...)

	for _, v := range ptxt {
		if b, e := k.BytesFor(v); e != nil {
			return nil, e
		} else if len(b) > 0 {
			ctxt = append(ctxt, b)
		}
	}

	return bytes.Join(ctxt, []byte(k.Mode.Divider())), nil
}

// EncryptFile will open the provided filename, read the contents, and
// encrypt using the provided Key.
func EncryptFile(fn string, k *Key) ([]byte, error) {
	var b []byte
	var e error

	if b, e = os.ReadFile(pathname.ExpandPath(fn)); e != nil {
		return nil, errors.Newf("failed to encrypt: %w", e)
	}

	return Encrypt(b, k)
}
