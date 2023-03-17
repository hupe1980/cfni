package cfni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXOR(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyz123456789"
	key := "abc4711"

	assert.Equal(t, text, xor(xor(text, key), key))
}
