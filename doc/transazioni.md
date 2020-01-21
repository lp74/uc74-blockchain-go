# Transazioni

## UTXO
nel modello Bitcoin il soggetto che utilizza un utxo in una nuova transazione deve:

- referenziarlo in ingresso fromando un `CTxIn` (referenza è la coppia `COutPoint`),
- fornire lo `scriptSig` (qui al momento si usa `PubKey`) per risolvere lo `scriptPubKey` dell'UTXO.

Nella forma più semplice dello `scriptPubKey` (Pay-to-Public-Key-Hash: P2PKH) lo `scriptSig` si compone di due campi:
- la firma di tutta la transazione <sig> (in verità una forma semplificata della transazione).
- la chiave pubblica del destinatario dell'UTXO <pubKey> con la quale verificare la firma <sig>

poiché:
- la validità dell'UTXO può essere verificata, (è appartenente ad una transazione di un blocco valido e non è stato mai speso) 
- e l'UTXO ha un `scriptPubKey` (qui `PubKeyHash`) che "comunica il legittimo proprietario":

fornendo lo `scriptSig` (qui `PubKey`) in `CTxIn` e firmando tutta la transazione con la chiave Privata `privateKey` 
è possibile verificare che il destinatario dell'UTXO è l'unico in grado di fornire una firma valida (attraverso la chiave privata) associata alla chiave pubblica dell'UTXO.

Nel nostro codice la firma di tutta la transazione `txn.Sign(privKey)`  va posta su tutti i `CTxIn` che compongono i `Vin` della transazione.