# Patterns

Here are examples of concurrency patterns.

## Publish Subscriber

This pattern involves three components:

* a publisher responsible for sending messages;
* a subscribe responsible for receiving messages;
* a broker to mediate between publishers and subscribers

This [working example](../cmd/pubsub/main.go) demonstrate this use of Go channels in publisher and subscriber pattern.

## Worker Pools

In other languages, a worker pool is intended to avoid spinning up unnecessary threads to process job. This does not apply to Go as the cost of spinning up Goroutine is minimal.

In Go worker pools are used to limit the number of concurrent processes at any one time to avoid overloading resources e.g. CPU core.

Here is a [working example](../cmd/workers/main.go) of a worker pool.