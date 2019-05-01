package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb3 "protoc/core"
	pb2 "protoc/discovery"
	pb "protoc/gossip"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	peerPort         = ":50051"
	gossipAddress    = "10.0.2.15"
	gossipPort       = ":50052"
	discoveryAddress = "10.0.2.15"
	discoveryPort    = ":50050"
)

var discoveryClient pb2.RegistryClient
var discoveryConn *grpc.ClientConn

//Peer stores the info regarding peer connection
type Peer struct {
	peerClient pb3.PeerManagerClient
	peerConn   *grpc.ClientConn
}

var peerList = make(map[string]Peer, 0)

type server struct{}

func (s *server) Deliver(ctx context.Context, in *pb.Block) (*empty.Empty, error) {
	var err error
	registerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// t := time.Now().String
	z, err := discoveryClient.FetchServiceLocation(registerCtx, &empty.Empty{}, grpc.FailFast(false))
	// connectPeer(z.Registrations[0])
	for i := 0; i < len(z.Registrations); i++ {
		if _, ok := peerList[(z.Registrations[i].Ipv4)]; ok {
			peerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			tempClient := peerList[(z.Registrations[i].Ipv4)]
			_, err = tempClient.peerClient.ValidateBlock(peerCtx, in, grpc.FailFast(false))
			cancel()
		} else {
			connectPeer(z.Registrations[i])
			peerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			fmt.Println("Reaches here")
			tempClient := peerList[(z.Registrations[i].Ipv4)]
			_, err = tempClient.peerClient.ValidateBlock(peerCtx, in, grpc.FailFast(false))
			fmt.Println("Reaches here2")
			cancel()
		}
	}

	if err != nil {
		log.Fatalf("could not greet: %v", err)
		fmt.Println(z)
	}

	return &empty.Empty{}, nil
}

func connectPeer(peerInfo *pb2.Registration) {
	var err error
	tempConn, err := grpc.Dial((peerInfo.Ipv4 + peerInfo.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	tempClient := pb3.NewPeerManagerClient(tempConn)
	peerList[peerInfo.Ipv4] = Peer{
		peerConn:   tempConn,
		peerClient: tempClient,
	}

}

func createDiscoveryConnection() {
	var err error
	discoveryConn, err = grpc.Dial((discoveryAddress + discoveryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	discoveryClient = pb2.NewRegistryClient(discoveryConn)
}

func closeDiscoveryConnection() {
	discoveryConn.Close()
}

func main() {

	go createDiscoveryConnection()

	lis, err := net.Listen("tcp", gossipPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("starting gossip")
	pb.RegisterGossipServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
