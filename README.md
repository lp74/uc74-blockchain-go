# Block Chain Go

Questo repository contiene il **codice** di una **Block Chain** costruita con **Go** a scopo **didattico**.

L'autore ha intenzionalmente scelto di scrivere i commenti in lingua italiana
per rendere più facile la lettura e la compresione del codice.
Anche se molti italiani hanno una buona conoscenza dell'inglese,
l'avere del materiale scritto in lingua madre può costituire una barriera in meno verso la comprensione.

> Scrivere una Block Chain didattica è un esercizio molto utile per capire più in dettaglio come funziona Bitcoin.
> Naturalemnte, in questa fase, il progetto è **molto semplice** e manca di molte funzionalità e sicurezze.

Pertanto mi scuso per le semplificazioni, le imprecisioni, gli eventuali errori.

**Segnalazioni e/o suggerimenti** saranno molto **apprezzati**.

## Guida al codice

La guida al codice è disponiile [qui](./doc/guida-al-codice.md).

## Piano di sviluppo

Ad oggi non posso stendere un piano di sviluppo.

Questo progetto lo seguo nel mio tempo libero (che sta diventando sempre meno)

Non so se potrò soddisfare questi desideri:

### fase 1

- [x] scrivere una semplice Block Chain con funzionalità di rete base

### fase 2 (IN CORSO)

- [ ] rifattorizzare il codice per ottenere una architettura migliore, seguendo i principi di design e i più comuni design pattern.
    - [ ] sviluppare una classe per la gestione della sicurezza (firma digitale)
    - [ ] sviluppare una classe per la gestione della testata del blocco
    - [ ] rifattorizzare la classe Block

- [ ] Migliorare i commenti e la documentazione, prendendo a riferimento Bitcoin

### fase 3

- [ ] estendere le funzionalità di networking per includere maggiori funzionalità peer-to-peer

### fase 4

- [ ] implementare (una parte) del protocollo Bitcoin

## Contribuire

Se desideri contribuire a questo progetto sei ben accetto.
Se sei uno studente, uno sviluppatore, o se semplicementi nutri interesse verso questo argomento
e desideri svilupparlo con me **non esitare a contattarmi**.

Sono disponibile per svolgere sessioni su skype e Hangouts. 

## Come compilare il codice e lanciare l'applicativo

Per lanciare l'applicazione digitare

```bash
go run main.go
```

Per compilare l'applicazione digitare

```bash
go build -o ./build/main main.go # seguire le istruzioni
```

## Ringraziamenti

Il codice prende spunto da tre progetti e da molto materiale web.
Il progetto è in corso. Sarà cura dell'autore accreditare i riferimenti di maggior rilievo. 

- [bitcoin](https://github.com/bitcoin/bitcoin)
- [btcsuite](https://github.com/btcsuite)
- [golang-blochain](https://github.com/tensor-programming/golang-blockchain)
