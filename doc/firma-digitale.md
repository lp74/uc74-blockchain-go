# Firma digitale


## TLDR;

<img src="./img/CPKDSS@2x.jpg" alt="bitcoin-address-generation" width="100%"/>

- generazione di una coppia chiave Privata, chiava Pubblica utilizzando ECDSA e secp256k1
- la chiave privata può firmare `sign()` un messaggio
- la chiave privata può verificare che il messaggio si stato firmato attraverso la chiave privata senza conoscerla `verify(pubKey, hashed_message, signature)`

Nel paragrafo [Transazioni](transazioni.md) viene descritto l'uso della firma digitale in Bitcoin.

# Generazione della chiave privata

Procedure per generare un coppia chiave privata, chiave pubblica:

1. inizializzare la curva ellittica. Nell'esempio si utilizza la curca P256 che è meno sicura della curva secp256k1.

2. utilizzando l'algoritmo ECDSA generate le chiavi. L'algoritmo va alimentato con la curva e un generatore di numero random crittograficamente sicuro.

3. la chiave pubblica viene generata concatendando i valori X e Y

## secp256k1

Bitcoin utilizza ECDSA e secp256k1.