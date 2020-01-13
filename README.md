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

