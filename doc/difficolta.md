# DIfficoltà

## Formato Compact

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
- se l'esponente è inferiore o uguale a 3 (non positivo):
  - translare a destra di 8*(exponent-3) bit la mantissa
- se l'esponente è maggiore di 3:
  - translare a sinistra di 8*(exponent-3) bit la mantissa
- applicare il segno

> In base 2 (binario) per computare M * 256^(E - 3) è sufficiente translate M a sinistra (o a destra se E - 3 è non positivo)
> di un numero di bit pari a 8 ∙ (E - 3).
> Infatti 256^( E - 3 ) = (2^8)^( E - 3 ) = 2^(8 ∙ ( E - 3 ))
