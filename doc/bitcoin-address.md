# Generate a Bitcoin-like address

Bitcoin utilizza una curva [ellittica Koblitz secp256k1](https://github.com/bitcoin-core/secp256k1) e l'algoritmo ECDSA.
La curva secp256k1 non era molto usata prima di bitcoin, ma adesso sta acquisendo popolarità a causa di alcune proprietà.
Molte delle curve utilizzate in crypto hanno una struttur random; secp256k1 è stata costruita in maniera non random che consente computazioni efficienti, risultando il 30% più veloce di altre curve. Inoltre è stata costruita in maniera tale da evitare l'inserimento di backdoor.

## Generate a new key pair

Procedure per generare un coppia chiave privata, chiave pubblica:

1. inizializzare la curva ellittica. Nell'esempio si utilizza la curca P256 che è meno sicura della curva secp256k1.
2. utilizzando l'algoritmo ECDSA generate le chiavi. L'algoritmo va alimentato con la curva e un generatore di numero random crittograficamente sicuro.
3. la chiave pubblica viene generata concatendando i valori X e Y