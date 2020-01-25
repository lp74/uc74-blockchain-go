# Blocco

Le informazioni che riguardano la struttura di un blocco sono contenute nei file bitcoin-core primitives/block.h e primitives/block.cpp.

Come si legge nel file, i nodi verificano e collezionano le transazioni dentro un blocco (con le logiche che vedremo) e ricavano un hash complessivo delle stesse attraverso una struttura dati che prende il nome di albero di Merkle.

Dunque risolvono l'algoritmo di consenso *Proof of Work* che consiste nella ricerca iterativa di un valore nonce `nNonce` da inserire nella testata del blocco (insieme alle altre informazioni) e che dia luogo a un hash di blocco inferiore ad un numero dato che prende il nome di *target*. Il target è collegato alla difficoltà `nBits`.

Risolto questo algoritmo trasferiscono il blocco agli altri nodi e tutti, se valido, lo inseriscono nella blockchain. La prima transazione del blocco è una transazione speciale che prende il nome di *coinbase*. Si tratta di una transazione che conia moneta a ricompensa del lavoro svolto il nodo che ha trasmesso per primo il nuovo blocco valido. Questa procedura è competitiva: tutti i nodi minatori costruiscoo il blocco destinando la transazione coinbase a un destinatario di loro scelta (in genere il minatore stesso).

> Dopo [SegWit](segregated-witness.md) l'albero di Merkle contiene anche script e firme tramite l'innesto di un sotto-albero per tramite della transazione coinbase.
