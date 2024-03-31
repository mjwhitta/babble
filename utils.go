package babble

import (
	urandom "crypto/rand"
	"math/big"
	"math/rand"
)

func randIntn(max int) int {
	var e error
	var n *big.Int

	if !CryptoSecure {
		return rand.Intn(max)
	}

	n, e = urandom.Int(urandom.Reader, big.NewInt(int64(max)))
	if e != nil {
		// Fallback to less secure PRNG
		return rand.Intn(max)
	}

	return int(n.Int64())
}
