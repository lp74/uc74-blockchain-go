# Canali

I canali sono condotti tipizzati attraverso i quali inviare e ricevere valori con l'operatore `<-`.

```go
ch <- v // invia v al canale ch
v := <- ch // riceve dal canale ch e assegna il valore a v.
```

I dati fluiscono nel verso della freccia.

Così come le mappe e le slices, i canali devono essere creati prima dell'uso.

```go
ch := make(chan int)
```

Di base, invia e riceve blocchi fino a quando l'altra parte è pronta. 
Questo consente la sincronizzazione di goroutine senza il lock specifico delle variabili.