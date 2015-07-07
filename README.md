# RabbitMQ

This repository implements part of the
[RabbitMQ Management HTTP API](https://cdn.rawgit.com/rabbitmq/rabbitmq-management/rabbitmq_v3_5_3/priv/www/api/index.html).

## Getting started

Go get the package:

```go
go get github.com/olivere/rabbitmq
```

Get an overview of the RabbitMQ cluster:

```go
client, err := rabbitmq.NewClient()
if err != nil {
	log.Fatal(err)
}

overview, err := client.Overview().Do()
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Running RabbitMQ %s", overview.RabbitMQVersion)
```

## Credits

We're all [standing on the shoulder of giants](https://en.wikipedia.org/wiki/Standing_on_the_shoulders_of_giants).

Thanks to the guys working on [RabbitMQ](https://www.rabbitmq.com)
and [Go](https://golang.org).

## License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
