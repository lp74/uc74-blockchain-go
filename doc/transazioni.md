# Transazioni

Il modello di transizione di cui si parla in questo paragrafo è antecedente a [SegWit](./segregated-witness.md) (*segregated witness*).

SegregatedWitness è stato introdotto con un dibattuto soft-fork all'altezza 481824 come è possibile vedere nel file chainparams.cpp di bitcoin-core.

```cpp
//...
consensus.SegwitHeight = 481824; // 0000000000000000001c8018d9cb3b742ef25114f27563e3fc4a1902167f9893
//...
```

Ciò non inficia la validità della trattazione in quanto le transazioni in questa forma sono ancora compatibili con il protocollo.

Nel proseguo del progetto verrà spiegato anche SegWit e vedremo quali sono le principali differenze.

<img src="./img/transaction.png" alt="transaction" width="50%"/>

## UTXO

Un UTXO è un output spendibile di una transazione valida. Nel modello Bitcoin, per spendere un UTXO occorre:

- referenziare l'UTXO in un ingresso `CTxIn` della transazione .
(la referenza è la coppia `COutPoint` composta dal riferimento alla transazione che continete l'UTXO e l'indice del medesimo all'interno di `Vout`),

- fornire lo `scriptSig` * che sblocca lo `scriptPubKey` dell'UTXO.

Nella forma più semplice di `scriptPubKey` (Pay-to-Public-Key-Hash: P2PKH) lo `scriptSig` si compone di due campi:

- la firma di tutta la transazione \<sig> (in verità una forma semplificata della transazione).
- la chiave pubblica del destinatario dell'UTXO \<pubKey> con la quale verificare la firma \<sig>

poiché:

- la validità dell'UTXO può essere verificata, (è appartenente ad una transazione di un blocco valido e non è stato mai speso) 
- e l'UTXO ha un `scriptPubKey` che "comunica il legittimo proprietario":

fornendo lo `scriptSig` in `CTxIn` e firmando la transazione con la chiave Privata `privateKey` è possibile verificare che il destinatario dell'UTXO è l'unico in grado di fornire una firma valida (attraverso la chiave privata) associata alla chiave pubblica dell'UTXO.

Nel nostro codice la firma  della transazione `txn.Sign(privKey)` va posta su tutti i `CTxIn` che compongono i `Vin` della transazione,
lo `scriptSig` è composto dalla coppia: (`PubKey`, `Signature`) e la `scripBubKey` è la `PubKeyHash`.
