package babble

import "regexp"

// Version is the package version.
const Version string = "0.3.3"

var (
	// CryptoSecure determines whether or not to use a
	// cryptographically secure PRNG.
	CryptoSecure bool = true

	footer       string         = "-----END BABBLE-----"
	header       string         = "-----BEGIN BABBLE-----"
	reWhiteSpace *regexp.Regexp = regexp.MustCompile(`\s+`)
)
