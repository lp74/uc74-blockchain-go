# La rete Bitcoin

## Architettura Peer-to-peer

Bitcoin è strutturato come una rete peer-to-peer (pari-a-pari) che si avvale di internet.
L'architettura peer-to-peer prevede che i singoli nodi (computer) siano collegati in una rete a maglia (mesh). Ciascuno nodo è pari agli altri e l'architettura è piatta. Tutti i nodi forniscono servizi e consumano servizi di rete. Non c'è un server centrale, non c'è gerarchia.

La decentralizzazione del controllo è uno degli elementi cardine del progetto Bitcoin e può essere ottenuta solo attraverso una architettura di questo tipo.

### Nodi e ruoli

Un nodo bitcoin è una collezione di funzioni:

* NETWORK: routing (instradamento)
* FULL BLOCKCHAIN: database (blockchain)
* MINER: mining (fabbricazione di nuovi blocchi)
* WALLET: wallet (gestione portafoglio)

Tutti i nodi includono la funzione di instradamento (NETWORK) e possono includere altre funzionalità.
Tutti i nodi validano e propagano transazioni, scoprono e mantengono le connessioni con altri peer.

Alcuni nodi (FULL BLOCKCHAIN) mantengono una copia aggiornata della blockchain e verificano le transazioni utilizzando un metodo chiamato Verifica Semplificata del Pagamento (SPV) 

Alcuni nodi (WALLET) non contengono una copia completa delle transazioni della blockchain.

I nodi (MINER) competono per creare nuovi blocchi. Richiedono hardware specializzato nella computazione di HASH (ad oggi ASIC). La creazione di un nuovo blocco che avviene attraverso la risoluzione competitiva di un "puzzle" connesso al contenuto del blocco stesso è incentivata ricompensando il nodo Miner vincente con una somma in BTC (ad oggi circa 12.5 BTC).

### Evoluzione architettura

* Bitcoin-core nodes N+W+B (non hanno funzionalità di mining dalla versione 0.13)
* Full Block Chain Node N+B
* Solo Miner M+B+N
* Lightweight (SPV) wallet N+W
* Pool Protocol Servers P, S (gateway per la connessione di altri nodi, come il pool di mining Stratum)
* Mining Nodes M+N
* Lightweight (SPV) W+D Stratum wallet

### Installare un Full Block Chain Node con Bitcoin-core

Installare un nodo Full è facile. Sul sito [BitcoinCore](https://bitcoin.org/en/bitcoin-core/) è possibile scaricare il pacchetto di installazione per il proprio sistema.

### lilp2p

[Libp2p](https://libp2p.io/) è una libreria che si compone di molti moduli e differenti parti:

* Transporto
* Stream muxers
* Canali criptati
* Connessioni e aggiornamento delle connessioni
* Instaradamento peer
* Salvataggio record
* Attraversamento NAT
* Ricerca (Discovery)
* Utilità di utilizzo generale e tipi
* Altre

Utilizzeremo questa libreria per fornire una architettura peer-to-peer alla nostra Block Chain didattica. Per farlo abbiamo bisogno di comprendere come funziona Bitcoin. Nel prossimo paragrafo vedremo come comunica i nodi fra di loro.

### Implementazione btcd

Btcd implementa l'architettura peer-to-peer (peer) e il protocollo bitcoin (wire).
Essendo scritta in Go può esserci di aiuto per comprendere come scrivere il nostro codice.

I pacchetti di rete sono:

* addrmgr
* netsync
* peer
* wire

### Come avviene la comunicazione fra nodi in Bitcoin
