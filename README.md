# UC74 Blockchain Go

## Introduzione alla tecnologia blockchain

Questo repository contiene il codice di una blockchain costruita a scopo didattico.

Per lanciare l'applicazione digitare

```bash
go run main.go
```

Per compilare l'applicazione digitare

```bash
go build -o ./build/main main.go
```

## Parte 3

### Transazioni

Fino ad adesso potevamo inserire qualsiasi dato nella catena, ma adesso viene il momento di inserire la cosa più importante di Bitcoin, le transazioni.
Una valuta digitale è una catena di firme digitali.
Chi detiene moneta può trasferirne una certa quantità al successivo firmando digitalmente la combinazione dell'hash di una transazione precedente con la chiave pubblica del proprietario successivo, aggiungendole alla fine della catena.
Colui che riceve un pagamento può verificare la catena di proprietà.

```bash
go build -o ./build/main main.go

# 18v37nkNq2cCnaZ4Z15LsAxhE3mGPcJwm1
# 1Pm2aZNxLbKmWMhDzmqjiWw5cEZ67UiyGz

./build/main createblockchain -address 1PsqSUvnn4Lkq5QaEaTMipUzpGiaS6H1Au

./build/main printchain

./build/main send -amount 1 -from 1PsqSUvnn4Lkq5QaEaTMipUzpGiaS6H1Au -to 1KQuw1NJ49eqZ9kE36mwGq5Jy8GoSBxx8w

```