package consensus

import "math/big"

// Params parametri che influenzano il consenso sulla validità della catena
type Params struct {
	/** Parametri Proof of work */
	PowLimit                *big.Int // il target più basso ammissibile per questa catena
	PowTargetSpacing        uint64
	PowTargetTimespan       uint64
	MinerConfirmationWindow uint64
	MinimumChainWork        *big.Int
}
