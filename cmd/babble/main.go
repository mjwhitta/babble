package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/mjwhitta/babble"
	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/log"
)

func debug(k *babble.Key) {
	var after string
	var b []byte
	var before string
	var e error
	var tmp [md5.Size]byte

	if b, e = os.ReadFile(cli.Arg(0)); e != nil {
		panic(e)
	}

	tmp = md5.Sum(b)
	before = hex.EncodeToString(tmp[:])
	log.Debugf("Original:  %s, %0.2fMB", before, mb(b))

	if b, e = babble.Encrypt(b, k); e != nil {
		panic(e)
	}

	tmp = md5.Sum(b)
	after = hex.EncodeToString(tmp[:])
	log.Debugf("Encrypted: %s, %0.2fMB", after, mb(b))

	k.Mode.Skip(0)
	if b, e = babble.Decrypt(b, k); e != nil {
		panic(e)
	}

	tmp = md5.Sum(b)
	after = hex.EncodeToString(tmp[:])
	log.Debugf("Decrypted: %s, %0.2fMB", after, mb(b))

	if before == after {
		log.Good("Success")
	} else {
		log.Err("Fail")
	}
}

func decrypt(k *babble.Key) {
	var b []byte
	var e error

	if b, e = babble.DecryptFile(cli.Arg(0), k); e != nil {
		panic(e)
	}

	if !flags.quiet {
		hl.Printf("%s", b)
		println()
	}
}

func encrypt(k *babble.Key) {
	var b []byte
	var e error

	if b, e = babble.EncryptFile(cli.Arg(0), k); e != nil {
		panic(e)
	}

	if !flags.quiet {
		switch flags.mode {
		case "byte":
			hl.Printf("%s", b)
		case "paragraph":
			hl.Printf("%s\n", b)
		default:
			hl.PrintfWrap(70, "%s\n", b)
		}
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			log.ErrX(Exception, r.(error).Error())
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

	mode.Skip(flags.skip)

	k, e = babble.NewKeyFromFile(flags.key, mode, flags.width)
	if e != nil {
		panic(e)
	}

	// Important b/c we don't want to skip content when parsing
	// ciphertext/plaintext
	mode.Skip(0)

	if cli.NArg() == 0 {
		hl.Println(k.String())
	} else {
		if flags.debug {
			debug(k)
		} else if flags.decrypt {
			decrypt(k)
		} else {
			encrypt(k)
		}
	}
}

func mb(b []byte) float64 {
	return float64(len(b)) / (1024 * 1024)
}
