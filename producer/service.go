package producer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type HairyLemon struct {
	UniqueId string  `json:"UniqueId"`
	FirstName    string  `json:"FirstName"`
	LastName    string  `json:"LastName"`
	Score   float64 `json:"Score"`
}

var DB []HairyLemon
var response []HairyLemon

// this function acts as a db to hold values for each 'unique id'
func responsevalues() []HairyLemon {
	DB = []HairyLemon{
		{"1", "Liam", "Gallagher", 7000.00},
		{"2", "Noel", "Gallagher", 2200.00},
		{"3", "Paul", "Arthurs", 3244.21},
		{"4", "Tony", "McCarroll", 11432.00},
		{"5", "Paul", "McGuigan", 5535.00}
	}
	return DB
}

var storevalues HairyLemon

type Kafka struct {
	AsyncProducer sarama.AsyncProducer
}

var Brokers = []string{"localhost:9092"}

// handler gets message and send it to the kafka queue asynchronously
func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("The key is: ", key)
	kafka := Kafka{
		AsyncProducer: NewAsyncProducer(Brokers),
	}
	for _, j := range responsevalues() {
		if j.UniqueId == key {
			storevalues = j
			fmt.Println(storevalues)
			prettyJSON, err := json.MarshalIndent(j, "", "    ")
			if err != nil {
				panic(err)
			}
			_, _ = fmt.Fprint(w, string(prettyJSON))

			kafka.AsyncProducer.Input() <- &sarama.ProducerMessage{
				Topic: "test_topic",
				Key:   sarama.StringEncoder(j.UniqueId),
				Value: sarama.StringEncoder(prettyJSON),
			}
		}
	}
}

func NewAsyncProducer(brokers []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy

	// wait until all pending messages are persisted and then reply success to the client
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	return producer
}
