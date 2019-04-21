package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "protoc/discovery"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

var peerList = make(map[string](*pb.Registration), 0)

const (
	port string = ":50050"
)

type server struct{}

//Register is used to implement registry.Register
func (s *server) Register(ctx context.Context, newPeer *pb.Registration) (*pb.RegisterResponse, error) {

	peerList[newPeer.Ipv4] = newPeer

	return &pb.RegisterResponse{Message: "registration successfull"}, nil
}

//Unregister is used to implement registry.Unregister
func (s *server) Unregister(ctx context.Context, peerInfo *pb.Registration) (*pb.RegisterResponse, error) {

	delete(peerList, peerInfo.Ipv4)
	return &pb.RegisterResponse{Message: "Unregistration successfull"}, nil
}

//FetchServiceLocation is useed to implement registry.FetchServiceLocation
func (s *server) FetchServiceLocation(ctx context.Context, in *empty.Empty) (*pb.RegistrationList, error) {

	var peers []*pb.Registration
	for _, value := range peerList {
		peers = append(peers, value)
	}
	return &pb.RegistrationList{Registrations: peers}, nil

}

func main() {
	fmt.Println("starting discovery")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("starting discovery")
	pb.RegisterRegistryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
