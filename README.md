# hairy-lemon

`hairy-lemon` is an implementation of a Kafka producer application.

## Goals

 1. Exposes API to some values to the user by taking unique id as an input value
 2. DB mockup for storing and returning values to the user when queried with a unique id
 3. Implements a Kafka producer to publish a successful query

To run the full capability, you'll need both the producer [hairy-lemon](https://github.com/lhmzhou/hairy-lemon) and the consumer [copper-face-jacks](https://github.com/lhmzhou/copper-face-jacks). `hairy-lemon` is only 50% of the capability.

## System Architecture

Below is the bird's eye view of the e2e initiative. The red enclosure represents `hairy-lemon`.

![alt diagram](https://github.com/lhmzhou/hairy-lemon/blob/master/image/hairy-lemon_arch.png)

## Prerequisites

- go v1.13.1+
- kafka container [shopify/kafka](https://hub.docker.com/r/spotify/kafka/)
- [dep](https://github.com/golang/dep) for managing dependencies

## Build

1. Clone [hairy-lemon](https://github.com/lhmzhou/hairy-lemon).

2. Navigate to the `hairy-lemon` folder and run the program in your terminal by executing:

     `go run hairy-lemon`

3. To initialize a kafka instance:

    a). Copy the [yaml file](https://github.com/confluentinc/examples/blob/5.3.1-post/cp-all-in-one/docker-compose.yml) to desktop. Open yaml file, and delete everything apart from `zookeeper` and `broker` configs. Save the file. Next, execute below commands in cli:

        $ docker-compose up
        $ docker ps


    b). Run below command to enter the container:

        $ docker exec -it testdocker /bin/bash


    c). To get into the kafka consumer and initialize it, run the following:

        $ cd /opt/kafka_2.11-0.10.1.0/bin/
        $ ./kafka-console-consumer.sh --bootstrap-server
        $ localhost:9092 --topic newtopic --from-beginning

     Note: Running Step 3c will display the messages in kafka and test.

4. `hairy-lemon` is launched and the kafka consumer is initialized. Proceed to open browser and hit the url at `http://localhost:8081/score/1`

You will see a json response on the webpage. The same response can be seen in the kafka container. Effectively, `UniqueId = 1` is queried on the DB via API. The corresponding values are given in the response. These values are then published to the kafka topic, and you can view them on the kafka consumer terminal accordingly.

## h/t

Maximum respect and many thanks to the developers on these open-source projects for making `hairy-lemon` possible:

[rcrowley/go-metrics](https://github.com/rcrowley/go-metrics)
</br>
[jcmturner/gofork](https://github.com/jcmturner/gofork)
</br>
[klauspost/compress](https://github.com/klauspost/compress)
</br>
[go-redis/redis](https://github.com/go-redis/redis)
</br>
[golang/snappy](https://github.com/golang/snappy)
</br>
[hashicorp/go-uuid](https://github.com/hashicorp/go-uuid)
</br>
[pierrec/lz4](https://github.com/pierrec/lz4)
</br>
[eapache @ Shopify](https://github.com/Shopify/sarama)
</br>
[davecgh/go-spew](https://github.com/davecgh/go-spew)
