# Custom Kafka-partitioner in GoLang

Atm this is a simple kind-of Hard-coded custom partitioner for my special use case, to forward some type of message to a special partition in my cluster.
This can be further improved with a better idea and do more fun stuff with it.

Run the Kafka cluster (In case you have a K8s Cluster read the code and change what's needed to be changed)
```bash
docker compose up -d
````

Run the producer to publish some messages:
```bash
go run producer/main.go
```

Run the consumer (Atm I'm just consuming my special partition but be my guest and serve yourselves with what you need)
```bash
go run consumer/main.go
```

This has no tests atm, `"Regression testing"? What's that? If it compiles, it is good; if it boots up, it is perfect` Linus Torvalds.
