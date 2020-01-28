# Hash

La funzione di hash è una funzione non invertibile che mappa dati di lunghezza arbitraria in valori di lunghezza predefinita.
Il dato sui quali si applica la funzione prende il nome di pre-immagine, il risultato della funzione, detto hash (o valore di hash, codice di hash, digest) ne è l'immagine.
L'immagine (hash) è un valore che in rappresentazione binaria è dato da numero definito di bit (ad esempio, a seconda dell'algoritmo utilizzato, 128, 160, 256, etc ...).

In crittografia la funzione di hash deve soddisfare, fra le altre, le seguenti proprietà:

- determinismo: fissando la pre-immagine la funzione di hash genera sempre la stessa immagine,
- efficienza: deve essere veloce da computare,
- resistenza alla pre-immagine, cioè sia computazionalmente intrattabile la ricerca di una stringa che dia come risultato l'immagine.
- resistenza alla seconda pre-immagine, ovvero che sia computazionalmente intrattabile la ricerca di una stringa il cui hash sia uguale a quello di una stringa data
- resistenza alle collisioni, è richiesto che sia computazionalmente intrattabile la ricerca di due stringhe fra loro diverse che diano luogo alla stessa immagine.
- effetto valanga: una piccola variazione della pre-immagine produce valori di hash molto diversi che appaiono incorrelati.

## Esempi

SHASUM256 e RIPEMD160 sono due funzioni di hash utilizzate in Bitcoin.
Nel contesto di applicazioni crittografiche chiamereo l'ingresso della funzione messaggio e lo indicheremo con la lettera M.

### effetto valanga

```bash
echo foo | time shasum -a 256
# output: b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c
# tempi: 0.04 real 0.02 user  0.01 sys
echo fou | time shasum -a 256
# 3fdf270878a29979f66bff2d055147240d449e627f3d908e61ddf9b638c24a2a
# tempi: 0.04 real 0.02 user  0.01 sys
```

## Alcune funzioni di hash notevoli In Bitcoin-core

### SHASUM256

SHASUM256 è un algoritmo di hash che mappa un messaggio in un valore hash composto da 256 bits (32 byte).
Nel codice bitcoin-core la classe responsabile per la generazione del valore di hash è CSHA256 (sha256.h)

L'ingresso è `const unsigned char* data`,
L'uscita è `const unsigned char hash[256]`

La classe CSHA256 è utilizzata dalla classe CHash256 che dato un `const unsigned char* data` produce un `const unsigned char hash[256]`
applicando due volte CSHA256. In altri termini l'ingresso viene trasformato applicando due volte SHASUM256.

Infine una funzione (overload) inline hash256 produce un blob opaco di 256 bit utilizzando CHash256 a partire da oggetti e vettori.

Semplificando:

hash = SHASUM256(SHASUM256(data))

Questo doppio hash shasum256 prende il nome di shasum256d ed è ritenuto più sicuro contro l'attacco "extension lenght" (si veda "Practical Cryptography by Ferguson and Schneier"):

[Why hashing twice?](https://crypto.stackexchange.com/questions/50017/why-hashing-twice)

Ciò nonostante va detto che l'argomento è controverso. Dunque non prenderemo, per il moento, posizione particolare.

[The puzzle of the double hash](https://medium.com/@craig_10243/the-puzzle-of-the-double-hash-968196edb06d)

### RIPEMD160

RIPEMD160 è un algoritmo di hash che mappa un messaggio in un valore hash composto da 160 bits (20 byte).
Nel codice bitcoin-core la classe responsabile per la generazione del valore di hash è CRIPEMD160 (ripemd160.h)

L'ingresso è `const unsigned char* data`,
L'uscita è `const unsigned char hash[160]`,

La classe CRIPEMD160 è utilizzata dalla classe CHash160 che dato un `const unsigned char* data` produce un `const unsigned char hash[160]`
applicando CRIPEMD160.

Infine una funzione (overload) inline hash160 produce un blob opaco di 160 bit utilizzando CHash160 a partire da oggetti e vettori.
