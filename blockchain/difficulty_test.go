package blockchain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompat(t *testing.T) {

	var in uint32
	in = 0x1b0404cb

	var out = CompactToBig(in)
	target := "404cb000000000000000000000000000000000000000000000000"
	expected := new(big.Int)
	expected.SetString(target, 16)
	assert.Equal(t, out, expected, fmt.Sprintf("output: %064x -> expected: %064x", out, expected), "nBits are equal")

}
