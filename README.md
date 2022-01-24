# kafka

An example of how to install and run Kafka locally.

---

### Install Java

Install Java from [here](https://www.oracle.com/java/technologies/downloads/).

You can verify it was installed correctly with `$ which java && java -version`.

---

### Download Kafka

Download the latest stable of version (3.1.0 at time of writing) of Kafka from [here](https://archive.apache.org/dist/kafka/).

Extract the tar files with `$ tar -xvzf kafka_2.13-3.1.0.tgz`. This will create a `kafka_2.13-3.1.0/` directory which should contain `bin/` and `config/` directories.

### Run Zookeeper

Zookeeper is an Apache service for distributed server coordination. Kafka requies a Zookeeper server to run. To start a Zookeeper server, `$ cd kafka_2.13-3.1.0/` and then run:
```
$ bin/zookeeper-server-start.sh config/zookeeper.properties
```

---

### Run Kafka Brokers

We'll use 3 Kafka brokers for this example. To do so, put the `config/server.{1-3}.properties` files from this repo in the Kafka `config` direcotry. These config files specify log file directories that must also be created, so do:
```
$ mkdir /tmp/kafka-logs1
$ mkdir /tmp/kafka-logs1
$ mkdir /tmp/kafka-logs1
```
Now start the Kafka brokers by running the following in separate terminals
```
$ bin/kafka-server-start.sh config/server.1.properties
$ bin/kafka-server-start.sh config/server.2.properties
$ bin/kafka-server-start.sh config/server.3.properties
```

---

### Create Topics

Kafka messages are send to topics, which must first be created. To create a topic for a service (e.g. `service1`) run:
```
$ bin/kafka-topics.sh --create --topic service1 --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2
```

This will create a topic called `service1`. `bootstrap-server` points to the address of any one of the Kafka brokers (doesn't matter which one). `paritions` is the number of brokers the topic data will be split between. `replication-factor` is the number of brokers each partition will be replicated on. You can list topics with:
```
bin/kafka-topics.sh --list --bootstrap-server localhost:9093
```
and get mroe info about a topic with:
```
bin/kafka-topics.sh --describe --topic service1 --bootstrap-server localhost:9093
```

---

### Using Kafka

You can test sending messages to a topic with:
```
bin/kafka-console-producer.sh --broker-list localhost:9093,localhost:9094,localhost:9095 --topic service1
```
and can test receiving messages from a topic with:
```
bin/kafka-console-consumer.sh --bootstrap-server localhost:9093 --topic service1
```
which will receive all messages sent to the topic from the time it starts. To see all messages in that topic, add the `--from-beginning` option.

This repo provides a very simple example of how Kafka could be used with the broker setup from above. In this case, there is an `audit` service that will be tracking two other services, `service1` and `service2`.

From each `service{1-2}/` directory, run `$ go run main.go`. This will start sending messages to Kafka topics, and you should see a stream of `wrote a message`s in standard output.

From the `audit/` direcoty, run `$ go run main.go`. This will start reading messages from Kafka topics, and you should see a stream of `received message from`s with message data in standard output. There may be a delay up to a minute before messages start being read from Kafka.
