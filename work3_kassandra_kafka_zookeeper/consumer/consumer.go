package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	sarama_cluster "github.com/bsm/sarama-cluster"
	"github.com/gocql/gocql"
)

const (
	SiteQuery string = `UPDATE site_stats SET
			views = views + ?,
			clicks = clicks + ?
			WHERE date = ?
			AND site_id = ?`
)

type Update struct {
	Date   string `json:"date"`
	SiteId int    `json:"site_id"`
	Views  int    `json:"views"`
	Clicks int    `json:"clicks"`
}

func main() {
	cassCluster := gocql.NewCluster("localhost:9042")
	cassCluster.Keyspace = "test"
	cassSession, err := cassCluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}
	defer cassSession.Close()

	kafConfig := sarama_cluster.NewConfig()
	kafConfig.ClientID = "test"
	kafConfig.Consumer.Return.Errors = true
	kafConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafConfig.Consumer.Offsets.CommitInterval = time.Second

	addrs := []string{"localhost:9092"}
	kafClient, err := sarama_cluster.NewClient(addrs, kafConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Validate Config.
	err = kafConfig.Validate()
	if err != nil {
		log.Fatal(err)
	}

	var topics = []string{"test"}
	consumer, err := sarama_cluster.NewConsumerFromClient(kafClient, "test", topics)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	for {
		select {
		case msg := <-consumer.Messages():
			consumer.MarkOffset(msg, "")
			go Process(msg, cassSession)
		case err := <-consumer.Errors():
			log.Println("Failed to consume message: ", err)
		}
	}
}

func Process(msg *sarama.ConsumerMessage, session *gocql.Session) {
	u := Update{}
	err := json.Unmarshal(msg.Value[:], &u)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	if err := session.Query(SiteQuery, 1, 1, "2021-08-06", 1).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
}
