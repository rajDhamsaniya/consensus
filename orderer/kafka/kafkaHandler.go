package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	pb2 "study/GitHub/consensus/protoc/gossip"

	"github.com/Shopify/sarama"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	gossipAddress = "10.0.2.15"
	gossipPort    = ":50052"
)

var producer sarama.AsyncProducer
var gossipClient pb2.GossipClient
var gossipConn *grpc.ClientConn

// StartConsumer for consume
func StartConsumer(produce sarama.AsyncProducer) {

	var err error

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
	bunch := make([]*pb2.EndorsedTx, 0)
	// timer := time.Now()
	// var txStock chan bool
	// fmt.Println("init", timer)
	ticker := time.NewTicker(1 * time.Second)
ConsumerLoop:
	for {

		select {
		case msg := <-partitionConsumer.Messages():
			if sarama.StringEncoder(msg.Key) == "Tx" {
				var tmp pb2.EndorsedTx
				err = proto.Unmarshal(msg.Value, &tmp)
				// fmt.Println("Tx  Here")
				tmp.TimeStamp = time.Now().String()
				// ticker = time.NewTicker(1 * time.Second)
				bunch = append(bunch, &tmp)
				// consumed++
				fmt.Println("tx added to bunch", len(bunch))
			} else {
				go cutBlock(bunch)
				fmt.Println("TTC block arrived")
				// timer = time.Now()
				ticker = time.NewTicker(1 * time.Second)
				bunch = bunch[:0]
			}
			if len(bunch) == 5 {
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

// GetProducer for getting producer
func GetProducer() sarama.AsyncProducer {
	return producer
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

// ConnectGossipService for connection to the service
func ConnectGossipService() {

	var err error
	gossipConn, err = grpc.Dial((gossipAddress + gossipPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	gossipClient = pb2.NewGossipClient(gossipConn)

}

// CloseGossipService for closing the connection
func CloseGossipService() {
	gossipConn.Close()
}

func cutBlock(bunch []*pb2.EndorsedTx) {
	for i := 0; i < len(bunch); i++ {
		fmt.Println("bunch prints: ", bunch[i].TimeStamp)
	}
	if gossipClient == nil {
		ConnectGossipService()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// t := time.Now().String
	z, err := gossipClient.Deliver(ctx, &pb2.Block{Sign: "sign", Bunch: bunch, OffSet: "asd"}, grpc.FailFast(false))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		fmt.Println(z)
	}
}
