package blockchain

// TxOutput si compone di un valore e della chiave pubblica del destinatario
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput in questa implementazione una transazione di input si compone di
// ID in riferimento a transazioni precedenti
// Out il riferimento alla transazione di uscita
// una stringa arbitraria
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock verify the signature
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked una transazione può essere sbloccata se la sua chiave pubblica corrisponde all'indirizzo dato
func (out *TxOutput) CanBeUnlocked(address string) bool {
	return out.PubKey == address
}
