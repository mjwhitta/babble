//nolint:godoclint // These are tests
package babble_test

import (
	urandom "crypto/rand"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/mjwhitta/babble"
	assert "github.com/stretchr/testify/require"
)

func aliceKey() []byte {
	var b []byte
	var e error

	b, e = os.ReadFile(filepath.Join("testdata", "alice.txt"))
	if e != nil {
		panic(e)
	}

	return b
}

func byteKey() []byte {
	var b []byte

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

	return b
}

func randomData(n int) []byte {
	var b []byte = make([]byte, n)

	_, _ = urandom.Read(b)

	return b
}

func TestMode(t *testing.T) {
	t.Parallel()

	type testData struct {
		key   []byte
		mode  babble.Mode
		name  string
		seek  int
		width int
	}

	var b []byte = randomData(16 * 1024) // 16KB
	var tests []testData = []testData{
		{byteKey(), &babble.ByteMode{}, "Bytes", 0, 4},
		{aliceKey(), &babble.ParagraphMode{}, "Paragraphs", 6, 3},
		{aliceKey(), &babble.SentenceMode{}, "Sentence", 2, 4},
		{aliceKey(), &babble.WordMode{}, "Words", 0, 5},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				var ctxt []byte
				var e error
				var k *babble.Key
				var ptxt []byte

				test.mode.Seek(test.seek)

				k, e = babble.NewKeyFromBytes(
					test.key,
					test.mode,
					test.width,
				)
				assert.NoError(t, e)
				assert.NotNil(t, k)

				test.mode.Seek(0)

				ctxt, e = babble.Encrypt(b, k)
				assert.NoError(t, e)

				test.mode.Seek(0)

				ptxt, e = babble.Decrypt(ctxt, k)
				assert.NoError(t, e)

				assert.Equal(t, b, ptxt)
			},
		)
	}
}
