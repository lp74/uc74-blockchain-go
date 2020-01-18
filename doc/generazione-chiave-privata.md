# Generazione della chiave privata

Procedure per generare un coppia chiave privata, chiave pubblica:

1. inizializzare la curva ellittica. Nell'esempio si utilizza la curca P256 che è meno sicura della curva secp256k1.

2. utilizzando l'algoritmo ECDSA generate le chiavi. L'algoritmo va alimentato con la curva e un generatore di numero random crittograficamente sicuro.

3. la chiave pubblica viene generata concatendando i valori X e Y