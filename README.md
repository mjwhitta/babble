# Babble

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/babble?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/babble)
![License](https://img.shields.io/github/license/mjwhitta/babble?style=for-the-badge)

## What is this?

Babble will use a provided key file to create a simple substitution
cipher in order to decrypt/encrypt files.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/babble
$ go install --ldflags "-s -w" --trimpath \
    github.com/mjwhitta/babble/cmd/babble@latest
```

Or compile from source:

```
$ git clone https://github.com/mjwhitta/babble.git
$ cd babble
$ git submodule update --init
$ make
```

## Usage

```
$ babble -k /path/to/key.txt /path/to/payload >payload.bab
$ babble -d -k /path/to/key.txt payload.bab >/path/to/payload
```

or create `main.go` similar to:

```
package main

import (
    _ "embed"
    "fmt"

    "github.com/mjwhitta/babble"
)

//go:embed payload.bab
var payload []byte

func main() {
    var b []byte
    var e error

    if b, e = babbleDecrypt(payload); e != nil {
        panic(e)
    }

    fmt.Printf("%s", b)
}
```

Then run the following commands:

```
$ babble -k /path/to/key.txt /path/to/payload >payload.bab
$ babble -k /path/to/key.txt >>main.go # This creates babbleDecrypt()
$ go run ./main.go
```

## Links

- [Source](https://github.com/mjwhitta/babble)
