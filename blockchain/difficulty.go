package blockchain

import (
	"math/big"
)

// CompactToBig restituisce il target in formato * big.Int
func CompactToBig(compact uint32) *big.Int {
	var result *big.Int

	mantissa := compact & 0x007fffff
	isNegative := compact&0x00800000 != 0
	exponent := uint(compact >> 24)

	if exponent <= 3 {
		mantissa >>= 8 * (3 - exponent)
		result = big.NewInt(int64(mantissa))
	} else {
		result = big.NewInt(int64(mantissa))
		result.Lsh(result, 8*(exponent-3))
	}

	if isNegative {
		result = result.Neg(result)
	}

	return result
}
