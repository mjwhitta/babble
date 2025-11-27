package babble

import (
	urandom "crypto/rand"
	"math/big"
	"math/rand/v2"
)

func randIntn(maxN int) int {
	var e error
	var n *big.Int

	if !CryptoSecure {
		//nolint:gosec // G404 - User's choice, not default
		return rand.IntN(maxN)
	}

	n, e = urandom.Int(urandom.Reader, big.NewInt(int64(maxN)))
	if e != nil {
		// Fallback to less secure PRNG
		//nolint:gosec // G404 - User's choice, not default
		return rand.IntN(maxN)
	}

	return int(n.Int64())
}
