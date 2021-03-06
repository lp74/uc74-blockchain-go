# Guida al codice

## Introduzione

Bitcoin è una valuta digitale **sperimentale** che consente pagamenti istantanei *border-less* (senza confini)
verso chiunque nel mondo e senza ricorso ad una autorità centrale (la banca).

Nel corso della storia moderna ci sono stati molti i tentativi di creare una valuta digitale,
ma fu Satoshi Nakamoto (uno pseudonimo dietro al quale si cela un lui, un lei, un gruppo di persone o una associazione)
che nel 2008 riuscì a risolvere il problema della doppia spesa (*double-spending*)
senza ricorrere alla "zecca" e dando vita alla tecnologia blockchain.

Le caratteristiche fondanti sono essenzialmente tre:

- architettura distribuita **peer-to-peer**
- ricorso alla **crittografia** (hashing e firma digitale)
- risoluzione del **Problema del Generale Bizantino** (*double-spending*) attraverso meccanismi di consenso (*Proof-of-Work* in Bitcoin)

### Architettura punto a punto

L'architettura punto-a-punto consente di distribuire il libro mastro (*ledger*) delle transazioni su tutti i nodi della rete,
indipendentemente da una organizzazione centrale. Per fare un parallelo, quando visitiamo un sito web richiediamo e riceviamo
contenuti da "un server centrale" (ad esempio Facebook o Google); nel caso punto-a-punto non esiste un server centrale
ma una collezione di nodi indipendenti che replicano i contenuti e che possono fornirli su richiesta.
Ne consegue che i contenuti non sono gestiti da un unico soggetto  ma sono custoditi, gestiti e resi disponibili da molti soggetti fra loro indipendenti.

### Crittografia

L'Hashing consente di associare dati di arbitraria lunghezza a codici univoci di lunghezza definita (e.g. 256 bit).
Possiamo pensare all'Hashing come una mappa unidirezionale che a partire da un contenuto ne deriva in maniera deterministica un indice univoco di lunghezza finita.
Il contenuto prende il nome di pre-immagine e l'hash prende il nome di immagine. Ad ogni pre-immagine corrisponde una ed una sola immagine (ai fini pratici).
L'algoritmo di hasing è veloce ma non è computazionalmente possibile la ricostruzione della pre-immagine a partire dal codice hash e la probabilità che due
contenuti diano luogo allo stesso hash, utilizzando algoritmi sicuri (tipo shasum-256) è praticamente irrilevante.
Inoltre, piccole variazioni del contenuto danno luogo a hash molto differenti.
Questo permette di referenziare utenti, transazioni e blocchi per mezzo di codici crittograficamente sicuri.

Ricapitolando, dato un messaggio, idealmente il suo hash:

- è univoco e deterministico dato un messaggio esiste uno ed un solo hash (in pratica)
- è unidirezionale, altamente improbabile risalire al messaggio originale
- è veloce da computare

La firma digitale (ECDSA in Bitcoin) consente di:

- dimostrare la proprietà di moneta (UTXO) attraverso la coppia chiave pubblica e privata
- firmare il trasferimento di moneta verso un secondo soggetto attraverso l'impiego della chiave privata
- consentire ai nodi partecipanti della rete la verifica della liceità delle operazioni

### Consenso

Il meccanismo di consenso permette ai nodi di concordare l'estensione del libro mastro inserendo nuove transazioni nella blockchain.
Il meccanismo di consenso di Bitcoin (*Proof of Work*) è progettato in modo tale da rendere praticamente impossibile
la modifica fraudolenta della catena. Perciò è intenzionalmente "difficile". L'uso del termine difficile non va confuso con il termine complicato.
*Proof of Work* è volutamente dispendioso in termini di tempo e risorse hardware e energetiche e fa parte degli stratagemmi progettuali atti a garantire la sicurezza della catena in un contesto distribuito.
Si può affermare che Bitcoin faccia sapiente uso della teoria dei giochi per garantire sicurezza ed equilibrio nel quadro di un insieme di nodi privi di fiducia reciproca, ove l'interesse individuale converge verso quello collettivo date le regole scritte nel codice stesso.
Il meccanismo di consenso è una di queste regole.

## Il codice contenuto in questo repositorio

Repositorio (repositorium o repostorium) è un termine latino :)

Questo progetto descrive le funzioni ed il protocollo ricorrendo al codice di una blockchain scritta in Go.
Al momento si tratta di un prototipo "didattico" efficace come introduzione all'argomento.

## Sommario

- [Hash](hash.md)
- [Firma Digitale](firma-digitale.md)
- [Indirizzi Bitcoin](indirizzi-bitcoin.md)
- [Transazioni](transazioni.md)
  - [Hash di una transazione](hash-transazione.md)
  - [Albero di Merkle](merkle.md)
- [Blocco](blocco.md)
  - [Hash di un blocco](hash-blocco.md)
- [Catena](chain.md)
- [UTXO](utxo.md)
- [Consenso](proof-of-work.md)
  - [Difficoltà](difficolta.md)
  - [Miner](miner.md)
  - [Sussidio](sussidio.md)
  - [Commissioni](fee.md)
- [Portafoglio](portafoglio.md)
- [Rete](network.md)
  - [DNS e indirizzi](dns.md)
  - [Procolollo](protocollo.md)
- [Glossario](glossario.md)
