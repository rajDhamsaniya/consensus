package kafka

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	pb "study/GitHub/consensus/protoc/orderer"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gogo/protobuf/proto"
)

var producer sarama.AsyncProducer

// StartConsumer for consume
func StartConsumer(produce sarama.AsyncProducer) {

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	// var producer sarama.AsyncProducer
	bunch := make([]pb.EndorsedTx, 0)
	// timer := time.Now()
	// var txStock chan bool
	// fmt.Println("init", timer)
	ticker := time.NewTicker(1 * time.Second)
ConsumerLoop:
	for {

		select {
		case msg := <-partitionConsumer.Messages():
			if sarama.StringEncoder(msg.Key) == "Tx" {
				var tmp pb.EndorsedTx
				err = proto.Unmarshal(msg.Value, &tmp)
				// fmt.Println("Tx  Here")
				tmp.TimeStamp = time.Now().String()
				// ticker = time.NewTicker(1 * time.Second)
				bunch = append(bunch, tmp)
				// consumed++
				fmt.Println("tx added to bunch", len(bunch))
			} else {
				go cutBlock(bunch)
				fmt.Println("TTC block arrived")
				// timer = time.Now()
				ticker = time.NewTicker(1 * time.Second)
				bunch = bunch[:0]
			}
			if len(bunch) == 3 {
				go cutBlock(bunch)
				fmt.Println("max blocklength reached")
				// timer = time.Now()
				ticker = time.NewTicker(1 * time.Second)
				bunch = bunch[:0]
				// fmt.Println("new Here and there", len(bunch))
			}
			// log.Println("Consumed message offset", sarama.StringEncoder(msg.Value))
			consumed++
			fmt.Println("msg", consumed, sarama.StringEncoder(msg.Key))
			// continue ConsumerLoop

		case <-signals:
			// fmt.Println("Its stupid")
			break ConsumerLoop

		case <-ticker.C:
			// fmt.Println("Ticker")
			if len(bunch) > 0 {
				message := &sarama.ProducerMessage{Topic: "test", Key: sarama.StringEncoder("TTC"), Value: sarama.StringEncoder("a")}
				producer.Input() <- message
				if hurrey := <-producer.Successes(); hurrey.Offset > 0 {
					// fmt.Println("")
				}
				ticker = time.NewTicker(1 * time.Second)
				// timer = time.Now()
				// continue ConsumerLoop
			} else {
				// timer = time.Now()
				ticker = time.NewTicker(5 * time.Second)
			}
		}
	}
}

// StartProducer for producer
func StartProducer() sarama.AsyncProducer {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	var err error
	producer, err = sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	return producer
}

// StopProducer for stoping
func StopProducer(producer sarama.AsyncProducer) {
	producer.AsyncClose()
}

// SendTx for ...
func SendTx(produce sarama.AsyncProducer, a []byte) {

	message := &sarama.ProducerMessage{Topic: "test", Key: sarama.StringEncoder("Tx"), Value: sarama.ByteEncoder(a)}
	// fmt.Println("Aye", message)
	producer.Input() <- message
	for {
		if hurrey := <-producer.Successes(); hurrey.Offset > 0 {
			// fmt.Println("qwertyuiopasdfghklzxcvbnm********************************************************************")
			return
		}
	}
}

func cutBlock(bunch []pb.EndorsedTx) {
	for i := 0; i < len(bunch); i++ {
		fmt.Println("bunch prints: ", bunch[i].TimeStamp)
	}

}
