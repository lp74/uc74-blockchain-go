package hash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	bv := []byte("abcde")
	var out = Hash(bv)
	target := "1d72b6eb7ba8b9709c790b33b40d8c46211958e13cf85dbcda0ed201a99f2fb9"
	assert.Equal(t, target, hex.EncodeToString(out))
}

func TestHashD(t *testing.T) {
	bv := []byte("abcde")
	var out = HashD(bv)
	target := "1d72b6eb7ba8b9709c790b33b40d8c46211958e13cf85dbcda0ed201a99f2fb9"
	assert.Equal(t, target, hex.EncodeToString(out))
}
