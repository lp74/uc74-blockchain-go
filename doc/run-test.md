# Come lanciare il progetto

## Scenario

Lo scenario implementato è il seguente:

- Il nodo centrale (FULL) crea la chain.
- Altri nodi (SPV) si collegano e scaricano la chain.
- Viene creato un nodo MINER che si connette al nodo centrale e scarica la catena.
- Il nodo SPV crea transazioni.
- Il nodo MINER riceve le transazioni e le mantiene in un mempool.
- Quando ci sono un numero sufficiente di transazioni il nodo MINER prepara un nuovo blocco.
- Quando il blocco è stato creato lo invia la nodo centrale.
- Il nodo SPV si sincronizza attraverso il nodo centrale.
- L'utente del nodo SPV controlla se la transazione è stata inserita nella chain.

Questo scenario è simile a Bitcoin. Anche se non abbiamo implementato una vera rete P2P abbiamo modo di comprendere i 
meccanismo fondamentali sottostanti Bitcoin.

## Compilazione

```bash
go build -o ./build/chain
```

## Lancio

Aprire 3 terminali

### Creazione INDIRIZZI E CHAIN

#### 1. Terminale FULL

```bash
export NODE_ID=3000

# Creare un SPV
go run main.go createwallet
# output:
# New address is: $FULL

# creare la chain
go run main.go createblockchain -address $FULL
# 2020/01/21 17:08:10 Replaying from value pointer: {Fid:0 Len:0 Offset:0}
# 2020/01/21 17:08:10 Iterating file id: 0
# 2020/01/21 17:08:10 Iteration took: 44.13µs
# 000a0d352c771254a22ea09b59e309e065e634dbe86d6f78b4eeaaba2cbf84f7
# Genesis created
# Finished!

# copiare il blocco di genesi
cp -R ./tmp/blocks_3000/ ./tmp/blocks_4000/
cp -R ./tmp/blocks_3000/ ./tmp/blocks_5000/
cp -R ./tmp/blocks_3000/ ./tmp/blocks_gen/
```

#### Terminale SPV

```bash
export NODE_ID=4000

# Creare un wallet
go run main.go createwallet
# output:
# New address is: $SPV
```

#### Terminale MINER

```bash
export NODE_ID=5000

# Creare un wallet
go run main.go createwallet
# output:
# New address is: $MINER
```

```bash
export FULL=1sQos3j68SmGChhDAB4C7W6m51kZj6CqbRt3D21gCt9tbDfRov
export SPV=12W2sTsanMjG9qGRXxuZi5veLbVm8w5dRgz3aGK6Lk271FF5ZXK
export MINER=12YwHafiCF7L3egfV6CAALkKdfxx9TXwGY8mRCxAM11xmrh9LY9

echo $FULL
echo $SPV
echo $MINER
```

### TRANSAZIONI e MINING

#### Terminale FULL (mine first transzion)

```bash
go run main.go send -amount 10 -from $FULL -to $MINER -mine
# 2020/01/21 17:16:05 Replaying from value pointer: {Fid:0 Len:42 Offset:1071}
# 2020/01/21 17:16:05 Iterating file id: 0
# 2020/01/21 17:16:05 Iteration took: 17.286µs
# 0009b59d9623212ab74d3308a6effe2c7fd842750ec65fc37bc0c2bafac5d860
# Success!

go run main.go startnode
```

#### Terminale SPVT (crea una transazione)

```bash
go run main.go startnode

# Received version command
# Received inv command
# Recevied inventory with 3 block
# Received block command
# Recevied a new block!
# Added block 00003564b5ede9ce13a1e9c8eca34680cecfda573bc26ce0a245179e4472c2bc
# Received block command
# Recevied a new block!
# Added block 0009b59d9623212ab74d3308a6effe2c7fd842750ec65fc37bc0c2bafac5d860
# Received block command
# Recevied a new block!
# Added block 000a0d352c771254a22ea09b59e309e065e634dbe86d6f78b4eeaaba2cbf84f7

# Uscire ctrl+c

go run main.go send -amount 1 -from $FULL -to $SPV
go run main.go send -amount 1 -from $FULL -to $SPV

go run main.go startnode
```

#### Terminale MINER (mining)

- quando un blocco è pronto lo invia al nodo centrale

```bash
go run main.go startnode -miner $MINER

# Received version command
# Received inv command
# Recevied inventory with 3 block
# Received block command
# Recevied a new block!
# Added block 00003564b5ede9ce13a1e9c8eca34680cecfda573bc26ce0a245179e4472c2bc
# Received block command
# Recevied a new block!
# Added block 0009b59d9623212ab74d3308a6effe2c7fd842750ec65fc37bc0c2bafac5d860
# Received block command
# Recevied a new block!
# Added block 000a0d352c771254a22ea09b59e309e065e634dbe86d6f78b4eeaaba2cbf84f7
```
