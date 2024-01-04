## Go Concurrency

Welcome to my brain dump covering topics related to Go concurrency.

<u>Data stream processing</u>

A fan out pattern is a concurrency pattern where a stream of data is fan out data to multiple goroutines for processing. 

<u>Worker pool pattern</u>

A worker pool pattern using a combination of goroutine and channel.

This working example [./example/summary/worker/main.go](../example/summary/worker/main.go) demonstrates the use of worker pool to build a fibonance sequence for a range of index using workers. To run the example, execute: `go run example/patterns/worker/main.go -seq=<sequence of fibonance> -workers=<number of workers>`.

<u>Observer pattern</u>

This working example [./example/summary/pubsub/main.go](../example/summary/pubsub/main.go) demonstrates an implementation of a pub sub (or listener) pattern using channels.
