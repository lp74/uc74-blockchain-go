# Come lanciare il progetto

## Compilazione

```bash
go build -o ./build/chain
```

## Lancio

Aprire 3 terminali

### Passo 1

#### 1. Terminale 1

```bash
export NODE_ID=3000

# Creare un wallet
go run main.go createwallet
# output:
# New address is: [ADDR_3000]

# creare la chain
go run main.go createblockchain -address [ADDR_3000]
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

#### 1. Terminale 2

```bash
export NODE_ID=4000

# Creare un wallet
go run main.go createwallet
# output:
# New address is: [ADDR_4000]
```

#### 1. Terminale 3

```bash
export NODE_ID=5000

# Creare un wallet
go run main.go createwallet
# output:
# New address is: [ADDR_5000]
```

### Passo 2

#### 2. Terminale 1: NODO FULL

```bash
go run main.go send -amount 10 -from [ADDR_3000] -to [ADDR_5000] -mine
# 2020/01/21 17:16:05 Replaying from value pointer: {Fid:0 Len:42 Offset:1071}
# 2020/01/21 17:16:05 Iterating file id: 0
# 2020/01/21 17:16:05 Iteration took: 17.286µs
# 0009b59d9623212ab74d3308a6effe2c7fd842750ec65fc37bc0c2bafac5d860
# Success!

go run main.go send -amount 1 -from [ADDR_3000] -to [ADDR_5000] -mine
# 2020/01/21 17:16:18 Replaying from value pointer: {Fid:0 Len:42 Offset:2803}
# 2020/01/21 17:16:18 Iterating file id: 0
# 2020/01/21 17:16:18 Iteration took: 22.635µs
# 00003564b5ede9ce13a1e9c8eca34680cecfda573bc26ce0a245179e4472c2bc
# Success!

go run main.go startnode
```

#### 2. Terminale 2: NODO SPV

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

go run main.go send -amount 1 -from [ADDR_5000] -to [ADDR_3000]

go run main.go startnode
```

#### 2. Terminale 3: NODO MINER

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

go run main.go startnode -miner [ADDR_5000]
```
