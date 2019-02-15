package main

import (
	pb "GitHub/grpc/discovery/registry"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var peers = make([](*pb.Registration), 0)

const (
	port string = ":50052"
)

type server struct{}

//Register is used to implement registry.Register
func (s *server) Register(ctx context.Context, newPeer *pb.Registration) (*pb.RegisterResponse, error) {

	peers = append(peers, newPeer)
	fmt.Println(peers)
	return &pb.RegisterResponse{Message: "registration successfull"}, nil
}

//Unregister is used to implement registry.Unregister
func (s *server) Unregister(ctx context.Context, peerInfo *pb.Registration) (*pb.RegisterResponse, error) {

	for i, pr := range peers {
		if pr == peerInfo {
			peers[i] = peers[len(peers)-1]
			peers = peers[:len(peers)-1]
			break
		}
	}
	return &pb.RegisterResponse{Message: "Unregistration successfull"}, nil
}

//FetchServiceLocation is useed to implement registry.FetchServiceLocation
func (s *server) FetchServiceLocation(ctx context.Context, rfr *pb.RegistrationFetchRequest) (*pb.RegistrationList, error) {
	// if rfr.FetchAll == true {
	return &pb.RegistrationList{Registrations: peers}, nil
	// }
	// else{
	// 	var req = make([](*pb.Registration), 0)

	// 	for pr := range peers {
	// 		if pr.name ==
	// 	}
	// }

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRegistryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
