# Difficoltà

## Premessa

La funzione di hash di Bitcoin mappa dati $ d \in \left{ \right} $

## Formato Compact nBits

In Bitcoin il campo nBits rappresenta la difficoltà per la Proof of Work di un dato blocco, da convertire nel valore target.

Il formato di rappresentazione è chiamato Compact e occupa 32 bit della testata.

Il codice di conversione può essere letto nel file arith_uint256.cpp.

Questa forma compatta è utilizzata per codificare numeri senza segno 256-bit che rappresentano il valore target ma implementano il segno per mantenersi consistenti con bitcond.

Si tratta di una rappresentazione simile ai numeri in virgola mobile IEEE754 che consiste di:

-------------------------------------------------
|   Esponente    |    Segno   |    Mantissa     |
| -------------- | ---------- | --------------- |
| 8 bits [31-24] | 1 bit [23] | 23 bits [22-00] |
-------------------------------------------------

Per passare alla rappresentazione come intero senza segno 32 bit occorre utilizzare la formula:

N = (-1^sign) ∙ M ∙ 256^(E-3)

Dato un valore in ingresso compact la a procedura è la seguente:

- ricavare la mantissa: compact & 0x007fffff (primi 23 bit)
- ricavare il segno: compact & 0x00800000 != 0
- ricavare l'esponente shiftando a destra 24 bit: uint(compact >> 24)
- se l'esponente è inferiore o uguale a 3 (non positivo) translare a destra di 8*(exponent-3) bit la mantissa
- se l'esponente è maggiore di 3 translare a sinistra di 8*(exponent-3) bit la mantissa
- applicare il segno

> In base 2 (binario) per computare M * 256^(E - 3) è sufficiente translate M a sinistra (o a destra se E - 3 è non positivo)
> di un numero di bit pari a 8 ∙ (E - 3).
> Infatti 256^( E - 3 ) = (2^8)^( E - 3 ) = 2^(8 ∙ ( E - 3 ))

## Formato leggibile

La procedura precedente permette di ricavare il valore decimale del target a partire dalla rappresentazione compact.
Il protocollo prevede la possibilità di calcolare la difficoltà `D` associata al valore `target`.
Per ricavare la difficoltà si utilizza la formula:

`D` = `TARGET_MAX` / `target`

Più nello specifico il codice è il seguente:

File: blockchain.cpp

```cpp
/* Calculate the difficulty for a given block index. */
double GetDifficulty(const CBlockIndex* blockindex)
{
    CHECK_NONFATAL(blockindex);

    int nShift = (blockindex->nBits >> 24) & 0xff;
    double dDiff = (double)0x0000ffff / (double)(blockindex->nBits & 0x00ffffff);

    while (nShift < 29)
    {
        dDiff *= 256.0;
        nShift++;
    }
    while (nShift > 29)
    {
        dDiff /= 256.0;
        nShift--;
    }

    return dDiff;
}
```

Per scelta progettuale il massimo valore target corrisponde a:

TARGET_MAX = 0x00000000FFFF0000000000000000000000000000000000000000000000000000
FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
Il target massimo ha 8 byte iniziali che corrispondono a 2^32 possibili valori, dunque a 2^32 hash.
Dividendo TARGET_MAX e il valore corrente del TARGET si ottiene la difficoltà che rappresenta l'aumeto relativo di difficoltà rispetto a quella di riferimento.
Ad esempio, in binario:

b00100 / b00010 = b10 = x2 = 2 = 2^1

b00100 / b00001 = b100 = x4 = 4 = 2^2

Dunque, a difficoltà doppia corrisponde uno spazio di ricerca di "un ordine di grandezza maggiore", seguendo legge espondenziale D = 2^k

Nel nostro esempio da 2*2 = 4 (b00100), si passa a 2^2 * 2^2 = 2*4 = 16 (b00001).

Questo è il motivo per cui il numero di hash da calcolare entro 600 secondi (10 minuti) è:

spazio di ricerca hashes = (D * 2^32) 600

Il rapporto l = ht / (2^32 * D) è il parametro di una distribuzione Poisson che descrive la probabilità che un blocco arrivi in 600 secondi.

## Adeguamente della difficoltà (del target)

La difficoltà non rimane costante nel tempo: ogni 2016 blocchi viene adeguata nel tenativo di mantenere un tempo medio di mining paria a 600 secondi per blocco.

Il protocollo opera nel modo seguente:

ogni 2016 blocchi  = 6 blocchi/ora * 24 ore/giorno * 14 giorni = 2016 blocchi


```cpp
// chainparams
consensus.nPowTargetTimespan = 14 * 24 * 60 * 60; // two weeks
consensus.nPowTargetSpacing = 10 * 60;

//

```cpp
// params.h
int64_t DifficultyAdjustmentInterval() const { return nPowTargetTimespan / nPowTargetSpacing; }
```

sottrae i timestamp dell'ultimo e del primo blocco. 
Idealmente questo valore dovrebbe essere (14 * 24 * 60 * 60) secondi, cioè due settimane,
Dunque questo valore viene diviso per 1209600. Se superiore a 4 o inferiore a -4 viene limitato all'intervallo [-4, 4].

Dunque la difficoltà del blocco successivo viene moltiplicata per questo valore.

la legge di crescita (decrescita) della difficoltà è dunque esponenziale. La difficoltà aumenta di 2^1 o 2^2.

Per questo motivo la curva di crescita H(t) corrisponde a spezzate logaritmiche.

Per maggiori dettagli:
- [Difficulty](https://en.bitcoin.it/wiki/Difficulty)
- [Block arrivals in the Bitcoin blockchain](https://arxiv.org/pdf/1801.07447.pdf)
