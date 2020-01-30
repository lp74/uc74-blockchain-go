package blockchain

import (
	"github.com/lp74/uc74-blockchain-go/consensus"
)

// CChainParams definisce i parametri di una istanza del sistema
type CChainParams struct {
	Consensus consensus.Params
}

// MainParams parametri della rete principale
func MainParams() (mainParams *CChainParams) {
	return &CChainParams{
		Consensus: consensus.Params{
			PowTargetTimespan: 4 * 24 * 60 * 60, // two weeks
			PowTargetSpacing:  10 * 60,
		},
	}
}
