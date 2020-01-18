# Block Chain Go

Questo repository contiene il **codice** di una **Block Chain** costruita con **Go** a scopo **didattico**.

L'autore ha intenzionalmente scelto di scrivere i commenti in lingua italiana,
per rendere più facile la lettura e la compresione del codice a sviluppatori italiani.
Anche se molti italiani hanno una buona conoscenza della lingua inglese,
l'avere del materiale scritto in lingua madre può costituire una barriera in meno verso la comprensione.

> Scrivere una Block Chain didattica è un esercizio molto utile per capire più in dettaglio come funziona Bitcoin.
> Naturalemnte, in questa fase, il progetto è molto semplice e manca di molte funzionalità e sicurezze.

Pertanto mi scuso per le semplificazioni, le imprecisioni, gli eventuali errori.

**Segnalazioni e/o suggerimenti** saranno molto **apprezzati**.

## Guida al codice

La guida al codice è disponiile [qui](./doc/guida-al-codice.md).

## Piano di sviluppo

Ad oggi non posso stendere un piano di sviluppo.

Questo progetto lo seguo nel mio tempo libero (che sta diventando sempre meno)

Non so se potrò soddisfare questi desideri:

### fase 1
- scrivere una semplice Block Chain con funzionalità di rete base

### fase 2
- Migliorare i commenti e la documentazione, prendendo a riferimento Bitcoin
- rifattorizzare il codice per ottenere una architettura migliore, seguendo i principi di design e i più comuni design pattern.

### fase 3
- estendere le funzionalità di networking fino a includere funzionalità peer-to-peer

### fase 4
- implementare (una parte) del protocollo Bitcoin

## Contribuire

Se desideri contribuire a questo progetto sei ben accetto


## Come compilare il codice e lanciare l'applicativo

Per lanciare l'applicazione digitare

```bash
go run main.go
```

Per compilare l'applicazione digitare

```bash
go build -o ./build/main main.go # seguire le istruzioni
```
