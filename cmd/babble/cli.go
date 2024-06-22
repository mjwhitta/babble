package main

import (
	"os"
	"strings"

	"github.com/mjwhitta/babble"
	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	debug   bool
	decrypt bool
	key     string
	mode    string
	nocolor bool
	quiet   bool
	secure  bool
	skip    int
	verbose bool
	version bool
	width   int
}

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf("%s [OPTIONS] [file]", os.Args[0])
	cli.BugEmail = "babble.bugs@whitta.dev"
	cli.ExitStatus(
		"Normally the exit status is 0. In the event of an error the",
		"exit status will be one of the below:\n\n",
		hl.Sprintf("%d: Invalid option\n", InvalidOption),
		hl.Sprintf("%d: Missing option\n", MissingOption),
		hl.Sprintf("%d: Invalid argument\n", InvalidArgument),
		hl.Sprintf("%d: Missing argument\n", MissingArgument),
		hl.Sprintf("%d: Extra argument\n", ExtraArgument),
		hl.Sprintf("%d: Exception", Exception),
	)
	cli.Info(
		"Babble will use the provided key file to create a simple",
		"substitution cipher in order to decrypt/encrypt files.",
	)
	cli.SectionAligned(
		"MODES",
		"|",
		"b, byte|Split on bytes.\n",
		"p, paragraph|Split on paragraphs.\n",
		"s, sentence|Split on sentences.\n",
		"w, word|Split on words.",
	)
	cli.Title = "Babble"

	// Parse cli flags
	cli.Flag(
		&flags.debug,
		"debug",
		false,
		"Test encrypt/decrypt/encrypt.",
		true,
	)
	cli.Flag(
		&flags.decrypt,
		"d",
		"decrypt",
		false,
		"Decrypt file (default: encrypt).",
	)
	cli.Flag(
		&flags.key,
		"k",
		"key",
		"",
		"File containing 256 unique tokens (mandatory).",
	)
	cli.Flag(
		&flags.mode,
		"m",
		"mode",
		"word",
		"Specify how to split the key file (default: word).",
	)
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(
		&flags.quiet,
		"q",
		"quiet",
		false,
		"No output other than errors.",
	)
	cli.Flag(
		&flags.secure,
		"secure",
		false,
		"Use cryptographically secure PRNG.",
	)
	cli.Flag(
		&flags.skip,
		"s",
		"skip",
		0,
		"Skip the first N tokens in key file (default: 0).",
	)
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show stacktrace, if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Flag(
		&flags.width,
		"w",
		"width",
		1,
		"Create mapping of byte to N tokens (default: 1).",
	)
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("babble version %s\n", babble.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() > 1 {
		cli.Usage(ExtraArgument)
	}

	// Key is mandatory
	if flags.key == "" {
		cli.Usage(MissingOption)
	}

	switch strings.ToLower(flags.mode) {
	case "b", "byte":
		flags.mode = "byte"
		flags.skip = 0
		flags.width = 1
	case "p", "paragraph":
		flags.mode = "paragraph"
	case "s", "sentence":
		flags.mode = "sentence"
	case "w", "word":
		flags.mode = "word"
	default:
		cli.Usage(InvalidOption)
	}

	babble.CryptoSecure = flags.secure

	if flags.skip < 0 {
		flags.skip = 0
	}

	if flags.width < 1 {
		flags.width = 1
	}
}
