# Transazioni

## UTXO
nel modello Bitcoin il soggetto che utilizza un utxo in una nuova transazione deve:

- referenziarlo in ingresso
- fornire lo scriptSig per sbloccarlo (qui al momento si usa PubKey)
- firmare la transazione.

poiché:
- la validità dell'utxo può essere verificata, (è appartenente ad una transazione di un blocco valido e non è stato mai speso) 
- l'utxo ha un scripSig (qui `PubKeyHash`) che comunica il legittimo proprietario
- fornendo la `PubKey` in `CTxIn` e firmando tutta la transazione con la chiave Privata `privKey`

è possibile verificare che il destinatario dell'utxo è l'unico in grado di fornire una firma valida (attraverso la chiave privata) associata alla chiave pubblica dell'utxo.

La firma di tutta la transazione `txn.Sign(privKey)`  va posta su tutti i `CTxIn` che compongono i `Vin` della transazione.