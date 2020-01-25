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

### Il ruolo dei nodi completi

I nodi completi (*Full Node*) si occupano di scaricare e verificare ogni blocco e ogni transazione **prima** di trasmetterlo al resto dei nodi.

I nodi archivio (*Archivial Node*) sono nodi completi che salvano su disco l'intera blockchain e che possono fornire dati storici (tuttua la catena) agli altri nodi.

I nodi sfrondati (*Pruned Nodes*) sono nodi completi che **non** salvano tutta la catena. Dopo un sufficiente numero di blocchi è possibile sfrondare i blocchi rimuovendo le transazioni. Infatti, per loro natura, le transazioni sono implicitamente contenute nella testata per mezzo dell'hash del nodo radice dell'albero di Merkle e gli output ancora spendibili sono gestiti e salvati in ua struttura dati diversa (UTXO).

Molti nodi leggeri (*SPV*) usanola rete e il protocollo Bitcoin per connettersi ai nodi completi.

Le regole di consenso non hanno necessariamente bisogno della rete Bitcoin e i nodi Miner possono ricorrere ad altre reti e protocolli per comunicare fra loro, come nel caso della rete ad alta velocità per la trasmissione dei blocchi (*high-speed block ralay network*).

### DNS

Quando un nodo parte per la prima volta non conosce gli indirizzi IP degli altri nodi.
Per scoprire gli indirizzi il nodo emette una (o più) richiesta DNS con il nome di un *DNS seed*.

La risposta alla richiesta contiene uno o più **DNS A records** completi con gli indirizzi IP dei nodi bitcoin che possono accettare connessioni in ingresso.

Ad esempio:

```bash
# Richiesta fatta con dig:
dig seed.bitcoin.sipa.be

# Risposta:
# ; <<>> DiG 9.10.6 <<>> seed.bitcoin.sipa.be
# ;; global options: +cmd
# ;; Got answer:
# ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 47156
# ;; flags: qr rd ra; QUERY: 1, ANSWER: 25, AUTHORITY: 1, ADDITIONAL: 1
# 
# ;; OPT PSEUDOSECTION:
# ; EDNS: version: 0, flags:; udp: 1280
# ;; QUESTION SECTION:
# ;seed.bitcoin.sipa.be.          IN      A
# 
# ;; ANSWER SECTION:
# seed.bitcoin.sipa.be.   749     IN      A       94.113.116.8
# seed.bitcoin.sipa.be.   749     IN      A       5.9.67.183
# seed.bitcoin.sipa.be.   749     IN      A       165.227.84.200
# [...]
```

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
