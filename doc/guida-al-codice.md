# Guida al codice

## Introduzione

Bitcoin è una valuta digitale **sperimentale** che consente pagamenti istantanei *boreder-less* (senza confini)
verso chiunque nel mondo e senza ricorso ad una autorità centrale (la banca).

Nel corso della storia moderna ci sono stati molti i tentativi di creare una valuta digitale,
ma fu Satoshi Nakamoto (uno pseudonimo dietro al quale si cela un lui, un lei, un gruppo di persone o una associazione)
che nel 2008 riuscì a risolvere il probela della doppia spesa (*double-spending*)
senza ricorrere alla "zecca", dando così vita alla tecnologia blockchain.

Le caratteristiche fondanti sono essenzialmente tre:

- architettura distribuita **peer-to-peer**
- ricorso alla **crittografia** (hashing e firma digitale)
- risoluzione del **Problema del Generale Bizantino** attraverso meccanismi di consenso (*Proof-of-Work* in Bitcoin)

L'architettura distribuita consente di distribuire il libro mastro (*ledger*) delle transazioni su tutti i nodi della rete

La firma digitale (ECDSA in Bitcoin) consente di:

- dimostrare la proprietà di moneta (UTXO) attraverso la coppia chiave pubblica e privata
- firmare il trasferimento di moneta verso un secondo soggetto attraverso l'impiengo della chiave privata
- consentire ai nodi partecipanti della rete la verifica della leicità delle operazioni

Il meccanismo di consenso consente di estendere il libro mastro inserendo nuove transazioni nella blockchain.
Il meccanismo di consenso di Bitcoin (*Proof of Work*) è progettato in modo tale da rendere praticamente impossibile 
la modifica fraudolenta della catena. Perciò è intenzionalmente "difficile", dove per difficoltà non si intende "complicato"
ma soltanto dispendisoso in termini di tempo e risorse energetiche.

Questo progetto descrive le funzioni in termini di codice e ambisce allo sviluppo di un prototipo "didattico".

## Sommario

- [Firma Digitale](firma-digitale.md)
- [Indirizzi Bitcoin](indirizzi-bitcoin.md)
- [Blocco](blocco.md)
  - [Hash di un blocco](hash-blocco.md)
- [Catena](chain.md)
- [Transazioni](transazioni.md)
  - [Hash di una transazione](hash-transazione.md)
  - [Albero di Merkle](merkle.md)
- [UTXO](utxo.md)
- [Consenso](proof-of-work.md)
  - [Miner](miner.md)
- [Rete](network.md)
- [Glossario](glossario.md)
