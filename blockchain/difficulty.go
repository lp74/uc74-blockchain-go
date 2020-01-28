package blockchain

import "math/big"

// CompactToBig restituisce il target in formato * big.Int
func CompactToBig(compact uint32) *big.Int {

	mantissa := compact & 0x007fffff
	isNegative := (compact & 0x00800000) != 0
	exponent := uint(compact >> 24)
	var bn *big.Int
	if exponent <= 3 {
		mantissa >>= 8 * (3 - exponent)
		bn = big.NewInt(int64(mantissa))
	} else {
		bn = big.NewInt(int64(mantissa))
		bn.Lsh(bn, 8*(exponent-3))
	}

	if isNegative {
		bn = bn.Neg(bn)
	}

	return bn
}
