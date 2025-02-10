package babble

import "regexp"

var (
	// CryptoSecure determines whether or not to use a
	// cryptographically secure PRNG.
	CryptoSecure bool
	footer       string         = "-----END BABBLE-----"
	header       string         = "-----BEGIN BABBLE-----"
	whiteSpace   *regexp.Regexp = regexp.MustCompile(`\s+`)
)

// Version is the package version.
const Version string = "0.2.2"
