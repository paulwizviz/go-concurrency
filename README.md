## Go Concurrency

Welcome to my brain dump covering Go concurrency patterns.

## Worker pool pattern

A worker pool pattern using a combination of goroutine and channel.

This [working example](./examples/worker/main.go) demonstrates the use of worker pool to build a fibonance sequence for a range of index using workers. To run the example, execute: `go run example/patterns/worker/main.go -seq=<sequence of fibonance> -workers=<number of workers>`.

## Observer pattern

This [working example](./examples/pubsub/main.go) demonstrates an implementation of a pub sub (or listener) pattern using channels.
