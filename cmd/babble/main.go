//nolint:wrapcheck // I'm not wrapping errors in a main package
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/mjwhitta/babble"
	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/log"
)

func debug(k *babble.Key) error {
	var after string
	var b []byte
	var before string
	var ctxt []byte
	var e error
	var n int = 32
	var ptxt []byte
	var tmp [sha512.Size]byte

	if b, e = os.ReadFile(cli.Arg(0)); e != nil {
		return e
	}

	tmp = sha512.Sum512(b)
	before = hex.EncodeToString(tmp[:])
	log.Debugf("Original:  %s (%0.2fMB)", before[0:n], mb(b))

	if ctxt, e = babble.Encrypt(b, k); e != nil {
		return e
	}

	tmp = sha512.Sum512(ctxt)
	after = hex.EncodeToString(tmp[:])
	log.Debugf("Encrypted: %s (%0.2fMB)", after[0:n], mb(ctxt))

	// Reset before decryption
	k.Mode.Seek(0)

	if flags.debug > 1 {
		ptxt, e = babble.DecryptCompare(ctxt, k, b)
	} else {
		ptxt, e = babble.Decrypt(ctxt, k)
	}

	if e != nil {
		return e
	}

	tmp = sha512.Sum512(ptxt)
	after = hex.EncodeToString(tmp[:])
	log.Debugf("Decrypted: %s (%0.2fMB)", after[0:n], mb(ptxt))

	if before == after {
		log.Good("Success")
	} else {
		log.Err("Fail")
	}

	return nil
}

func decrypt(k *babble.Key) error {
	var b []byte
	var e error

	if b, e = babble.DecryptFile(cli.Arg(0), k); e != nil {
		return e
	}

	if !flags.quiet {
		fmt.Printf("%s", b)
		println()
	}

	return nil
}

func encrypt(k *babble.Key) error {
	var b []byte
	var e error

	if b, e = babble.EncryptFile(cli.Arg(0), k); e != nil {
		return e
	}

	if !flags.quiet {
		switch flags.mode {
		case "byte":
			fmt.Printf("%s", b)
		case "paragraph":
			fmt.Printf("%s\n", b)
		default:
			//nolint:mnd // I like to wrap at 70
			fmt.Printf("%s\n", hl.Wrap(70, string(b)))
		}
	}

	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r)
			}

			switch r := r.(type) {
			case error:
				log.ErrX(Exception, r.Error())
			case string:
				log.ErrX(Exception, r)
			}
		}
	}()

	var e error
	var k *babble.Key
	var mode babble.Mode

	validate()

	switch flags.mode {
	case "byte":
		mode = &babble.ByteMode{}
	case "paragraph":
		mode = &babble.ParagraphMode{}
	case "sentence":
		mode = &babble.SentenceMode{}
	case "word":
		mode = &babble.WordMode{}
	}

	mode.Seek(flags.skip)

	k, e = babble.NewKeyFromFile(flags.key, mode, flags.width)
	if e != nil {
		panic(e)
	}

	// Reset before decryption/encryption
	mode.Seek(0)

	if cli.NArg() == 0 {
		fmt.Println(k.String())
	} else {
		switch {
		case flags.debug > 0:
			e = debug(k)
		case flags.decrypt:
			e = decrypt(k)
		default:
			e = encrypt(k)
		}

		if e != nil {
			panic(e)
		}
	}
}

func mb(b []byte) float64 {
	//nolint:mnd // !MB
	return float64(len(b)) / (1024 * 1024)
}
