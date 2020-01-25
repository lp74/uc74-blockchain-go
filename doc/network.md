# La rete Bitcoin

## Architettura Peer-to-peer

Bitcoin è strutturato come una rete peer-to-peer (pari-a-pari) che si avvale di internet.
L'architettura peer-to-peer prevede che i singoli nodi (computer) siano collegati in una rete a maglia (mesh). Ciascuno nodo è pari agli altri e l'architettura è piatta. Tutti i nodi forniscono servizi e consumano servizi di rete. Non c'è un server centrale, non c'è gerarchia.

La decentralizzazione del controllo è uno degli elementi cardine del progetto Bitcoin e può essere ottenuta solo attraverso una architettura di questo tipo.

### Nodi e ruoli

Un nodo bitcoin è una collezione di funzioni:

* NETWORK: routing (instradamento)
* FULL BLOCKCHAIN: database (blockchain) detto anche Validating node perchè valida tutte le regole.
* MINER: mining (fabbricazione di nuovi blocchi)
* WALLET: wallet (gestione portafoglio)

Tutti i nodi includono la funzione di instradamento (NETWORK) e possono includere altre funzionalità.
Tutti i nodi validano e propagano transazioni, scoprono e mantengono le connessioni con altri peer.

Alcuni nodi (FULL BLOCKCHAIN) mantengono una copia aggiornata della blockchain e verificano le transazioni utilizzando un metodo chiamato Verifica Semplificata del Pagamento (SPV).

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

## Cosa è una rete pari-a-pari (P2P)

* come si propagano transazioni e blocchi ai nodi Bitcoin
* aperta, piatta (*flat*), nessuna autenticazione, nessun nodo speciale
* deve resistere agli attacchi
  * Attacchi DoS *Denyal of Service*
  * Attacchi Sybil

## Quali sono i comandi di rete

* VERSION - version message
* VERACK - riconosce la versione di un messaggio precedente comunicando al nodo che può iniziare a inviare messaggi
* ADDR - trasmette informazioni sui nodi peers noti (indirizzi IP)
* INV - trasmette uno o più inventari degli oggetti conosciuti
* GETDATA - richiede uno o più dati a un altro nodo
* GETBLOCKS - richiede un INV con la lista delle testate a partire da un punto particolare
* GETHEADERS - richiede le testata a partire da un particolare punto
* GETADDR - richiede un messaggio ADDR
* TX - trasmette una singola transazione
* HEADERS - trasmette le testate in risposta a GETHEADERS
* BLOCK - trasmette la serializzazione di un singolo blocco
* PING - inviato periodicamente per comunicare al nodo peers che il nodo trasmittente è ancora in vita
* PONG - risponde al comando PING, fornendo il nodo che sta rispondendo

altri comandi:

* MERKLEBLOCK
* MEMPOOL
* NOTFOUND
* FILTERLOAD
* FILTERADD
* FILTERCLEAR
* SENDHEADERS
* FEEFILTER
* SENDCMPCT
* CMPCTBLOCK
* GETBLOCKTXN
* BLOCKTXN

la lista completa è visionabile nel file [protocol.h](https://github.com/bitcoin/bitcoin/blob/master/src/protocol.h)

## Connessione alla rete (P2P)

* inizialmente i nodi si connettono ad uno o più nodi
* i nodi acquisiscono gli indirizzi di altri nodi della rete utilizzando messaggi **ADDR** attraverso un meccanismo di *Gossip* (pettegolezzo)
* un nodo Bitcoin (nell'implementazione bitcoin core) si connette fino ad altri 8 nodi in modalità *outbound* (ricezione dati)
* I nodi possono o meno accettare connessioni *inbound* da altri nodi

## Disconnession e censura dei nodi (*Banning*)

i nodi che si comportano male (fraudolenti o spammers) devono essere rimossi:

* consumano risorse
* diminuiscono l'accesso a nodi onesti

Comportamenti malevoli includono:

* trasmissione di transazioni e blocchi invalidi
* trasmettere malevolmente blocchi non connessi
  la nostra catena potrebbe difettare di alcuni blocchi *tip* ma se continuano ad inviare malevolemnte blocchi disconessi vanno censurati
* che si bloccano o inviano informazioni in maniera troppo lenta
* inviano transazioni che non rispettano lo standard
* messaggi malformati (prima di CSV)

In funzione del comportamento, Bitcoin Core può:

* ignorare il problema e continuare
* disconnettere il node peer immediatamente
* censurare il nodo peer per 24 ore (basato su IP)
* applicare punti DoS. Quando il punteggio totalizza 100 censurare il node peer.

## Tipi di nodi

### Nodo completo *Full Node*

* viene chiamato anche nodo pienamente verificatore *fully validating*
* riceve i blocchi non appena minati e propagati attraverso la rete
* verifica la validità dei blocchi e delle transazioni che includono prima di propagarli sulla rete
* rafforza le regole di consenso della rete Bitcoin
* mantiene una collezione di UTXO
* è il modo più sicuro e privato di utilizzare Bitcoin, perché l'indirizzo o la chiave pubblica non vengono rilevati.

### Nodo d'archiviazione *Archivial Node*

I nodi archivio (*Archivial Node*) sono nodi completi che salvano su disco l'intera blockchain e che possono fornire dati storici (tutta la catena) agli altri nodi.

### Nodo sfrondato

* è un tipo di *Full Node*
* scarta i dati di blocchi vecchi per risparmiare spazio su disco. (Ad oggi mantenere un nodo completo richiede circa 300 GB di spazio su disco)
* ritiene gli ultimi due giorni in maniera completa e permette le riorganizzazioni.
* propaga i nuovi blocchi ma non può servire i vecchi
* è sicuro come un nodo completo

### Nodo per la verifica semplificata del pagamento (*Simple Payment Verification* - *SPV*)

I nodi sfrondati (*Pruned Nodes*) sono nodi completi che **non** salvano tutta la catena. Dopo un sufficiente numero di blocchi è possibile sfrondare i blocchi rimuovendo le transazioni. Infatti, per loro natura, le transazioni sono implicitamente contenute nella testata per mezzo dell'hash del nodo radice dell'albero di Merkle e gli output ancora spendibili sono gestiti e salvati in ua struttura dati diversa (UTXO).

* scarica soltanto:
  * la testata dei blocchi
  * informazioni circa transazioni specifiche
* possono validate la *Proof of Works* in quant serve solo la testata
* possono verificare che una transazione sia inclusa in un blocco (richiedendo la prova Merkle)
* possono usare i filtri bloom per preservare un pò di privacy
* *non possono* verificare la doppia spesa
* *non possono* verificare la disponibilità di moneta
* *non possono* verificare se una transazione è stata confermata dentro la blockchain

### Altre opzioni

* *-blocksonly* un nodo completo che non propaga informazioni
* *-nolisten* un nodo che produce connessioni in uscita ma non in ingresso
* *-onion* un nodo che si connette ai peers utilizzando tor
* *-proxy* si connette ai nodi tramite un proxy
* *-whitelist=\<IP address or subnet>* non censura mai i nodi sulla *whitelist*

Molti nodi *SPV* usano la rete e il protocollo Bitcoin per connettersi ai nodi completi.

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

### Riferimento al codice

In bitcoin-core i DNS seeds sono hardcoded nel file chainparams.cpp:

```cpp
// ...
// std::vector<std::string> vSeeds;
// ...
vSeeds.emplace_back("seed.bitcoin.sipa.be"); // Pieter Wuille, only supports x1, x5, x9, and xd
vSeeds.emplace_back("dnsseed.bluematt.me"); // Matt Corallo, only supports x9
vSeeds.emplace_back("dnsseed.bitcoin.dashjr.org"); // Luke Dashjr
vSeeds.emplace_back("seed.bitcoinstats.com"); // Christian Decker, supports x1 - xf
vSeeds.emplace_back("seed.bitcoin.jonasschnelli.ch"); // Jonas Schnelli, only supports x1, x5, x9, and xd
vSeeds.emplace_back("seed.btc.petertodd.org"); // Peter Todd, only supports x1, x5, x9, and xd
vSeeds.emplace_back("seed.bitcoin.sprovoost.nl"); // Sjors Provoost
vSeeds.emplace_back("dnsseed.emzy.de"); // Stephan Oeste
// ...
```

## Formato del messaggio

I messaggi Bitcoin P2P contengono una testata *header* e un carico *payload*

* La testata si compone di 24 byte:
  * Message start (4 byte): indicano la rete (Mainnet: 0xf9beb4d9)
  * Comand name (12 byte): eg ADDR, INV, BLOCK, etc ...
  * Payload size (4 bytes): comunica la dimensione del payload
  * Checksum (4 byte): doppio SHA256 del payload
* Il carico è fino a 32 MB e ogni comando ha il suo formato.

| MSG START | Command name | Payload size | Checksum | Body         |
| --------- | ------------ | ------------ | -------- | ------------ |
| 4 byte    | 12 byte      | 4 byte       | 4 byte   | fino a 32 MB |

```cpp
// chainparams.cpp

//...
pchMessageStart[0] = 0xf9;
pchMessageStart[1] = 0xbe;
pchMessageStart[2] = 0xb4;
pchMessageStart[3] = 0xd9;
//...
```

## Controllo messaggi

### Version handshake

... 19
