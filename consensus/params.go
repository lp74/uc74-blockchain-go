package consensus

import "math/big"

// Params parametri che influenzano il consenso sulla validità della catena
type Params struct {

	/** Parametri Proof of work */
	PowLimit *big.Int
}
