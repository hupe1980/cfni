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

func TestToPythonList(t *testing.T) {
	expected := `["a", "b", "c"]`
	input := []string{"a", "b", "c"}

	assert.Equal(t, expected, toPythonList(input))
}

func TestToPythonDict(t *testing.T) {
	expected := `{"a": "va"}`
	input := map[string]string{"a": "va"}

	assert.Equal(t, expected, toPythonDict(input))
}
