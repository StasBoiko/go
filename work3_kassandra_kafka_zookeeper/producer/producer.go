package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type Update struct {
	Date   string `json:"date"`
	SiteId int    `json:"site_id"`
	Views  int    `json:"views"`
	Clicks int    `json:"clicks"`
}

func main() {
	var producer sarama.SyncProducer
	brokers := []string{"localhost:9092"}
	сonfig := sarama.NewConfig()
	сonfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, сonfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to Kafka brokers: %s\n", err)
		os.Exit(1)
	}

	defer func() {
		producer.Close()
	}()

	m := Update{
		Date:   "2021-08-06",
		SiteId: 2,
		Views:  1,
		Clicks: 1,
	}
	jsonMsg, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}

	msg := sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.ByteEncoder(jsonMsg),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sent msg to partition:", partition, ", offset:", offset)
	}
}
