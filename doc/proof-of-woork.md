# Proof Of Work

Bitcoin-core esegue la PoW in mining.cpp all'interno della funzione generateBlocks.

```cpp
// permette di inizializzare il nuovo blocco prima della PoW con
void IncrementExtraNonce(CBlock* pblock, const CBlockIndex* pindexPrev, unsigned int& nExtraNonce)
{
  //...
}

```

Poi con un ciclo while viene modificato solo il nounce

```cpp

while (nMaxTries > 0 && pblock->nNonce < std::numeric_limits<uint32_t>::max() && !CheckProofOfWork(pblock->GetHash(), pblock->nBits, Params().GetConsensus()) && !ShutdownRequested()) {
    ++pblock->nNonce; // qui
    --nMaxTries;
}
```

GetNextWorkRequired è chiamata da UpdateTime chiamato da CreateNewBlock


In maniera analogo dobbiamo modificare la PoW per evitare di inizializzare i dati ad ogni step.
In analogia sposterò generateBlock in mining. Al momento MineTx in Network

## nBits

Il calcolo del TARGET da non confondere con la difficoltà è [qui](./difficolta).
[come calcolare nBits](https://bitcoin.stackexchange.com/questions/2924/how-to-calculate-new-bits-value)