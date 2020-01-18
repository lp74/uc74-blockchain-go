# Indirizzi Bitcoin

Bitcoin utilizza ECDSA (curva secp256k), SHA-256, RIPEMD-160 e la codifica Base58Check per generare un indirizzo.

L'immagine seguente mostra come ottenere un indirizzo a partire dalla chiave Pubblica (versione 1).

---

<img src="https://en.bitcoin.it/w/images/en/9/9b/PubKeyToAddr.png" alt="bitcoin-address-generation" width="50%"/>

---

## Procedura utilizzata

Vediamo dunque nel dettaglio quali sono i passi necessari per generare l'indirizzo a partire dalla chiave pubblica.

Nella nostra applicazione svolgiamo una procedura differente rispetto all'immagine.

Per la generazione della chiave privata e la derivazione della relativa chiave pubblica, riferirsi al paragrafo [generazione della chiave privata](generazione-chiave-privata)

1. avendo una chiave privata ECDSA, prelevare la corrispondente chiave pubblica.
La chiave pubblica è un punto sulla curva ellittica le cui cordinate sono (X,Y). Il nostro applicativo concatena X e Y per generare la chiave pubblica `publicKey`

2. computa l'hash SHA-256 della chiave pubblica `pubKeyHash`
la lunghezza dell'hash è 256 bits

3. computa l'hash RIPEMD-160 dell'hash della chiave pubblica `ripemd160`
la lunghezza dell'hash è 160 bits, ovvero 20 bytes
4. preprende il byte di versione `versionedRimpemd160`
5. computa il `checksum` del `versionedRimpemd160`
6. concatena il `versionedRimpemd160` e il `checksum` per formare `fullHash`
7. computa la codicifica Base58 di `fullHash` per generare `address`

## Note

Bitcoin utilizza una curva [ellittica Koblitz secp256k1](https://github.com/bitcoin-core/secp256k1) e l'algoritmo ECDSA.
La curva secp256k1 non era molto usata prima di bitcoin, ma adesso sta acquisendo popolarità a causa di alcune proprietà.
Molte delle curve utilizzate in crypto hanno una struttur random; secp256k1 è stata costruita in maniera non random che consente computazioni efficienti, risultando il 30% più veloce di altre curve. Inoltre è stata costruita in maniera tale da evitare l'inserimento di backdoor.
