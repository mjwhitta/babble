package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

var fn string

func init() {
	flag.Usage = func() {
		println("Usage: bytemap <file>")
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	fn = flag.Arg(0)
}

func main() {
	var b []byte
	var bytes []string
	var e error

	for i := 0; i < 256; i++ {
		bytes = append(bytes, fmt.Sprintf("%02x", i))
	}

	rand.Shuffle(
		256,
		func(i int, j int) {
			var tmp string = bytes[i]

			bytes[i] = bytes[j]
			bytes[j] = tmp
		},
	)

	if b, e = hex.DecodeString(strings.Join(bytes, "")); e != nil {
		println(e.Error())
	} else {
		if e = os.WriteFile(fn, b, 0o644); e != nil {
			println(e.Error())
		}
	}
}
