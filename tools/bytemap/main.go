package main

import (
	"flag"
	"math"
	"math/rand/v2"
	"os"
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
	var e error

	for i := range math.MaxUint8 + 1 {
		b = append(b, byte(i))
	}

	rand.Shuffle(
		math.MaxUint8+1,
		func(i int, j int) {
			var tmp byte = b[i]

			b[i] = b[j]
			b[j] = tmp
		},
	)

	//nolint:mnd // u=rw,go=-
	if e = os.WriteFile(fn, b, 0o600); e != nil {
		println(e.Error())
	}
}
