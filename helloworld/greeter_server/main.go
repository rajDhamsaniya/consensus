/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb3 "../../protoc/contractcode"
	pb2 "../../protoc/discovery"
	pb "../../protoc/helloworld"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	registryAddress = "10.0.2.15"
	contractAddress = "10.0.2.15"
	defaultName     = "10.0.2.15"
	peerPort        = ":50051"
	contractPort    = ":50053"
	registryPort    = ":50052"
	sign            = "Sign"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	time.Sleep(2 * time.Second)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

func registerService() {
	conn, err := grpc.Dial((registryAddress + registryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb2.NewRegistryClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.Register(ctx, &pb2.Registration{Name: name, Ipv4: defaultName, Port: peerPort}, grpc.FailFast(false))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func (s *server) ExecuteTransaction(ctx context.Context, in *pb.Executetx) (*pb.ExecResponse, error) {

	conn, err := grpc.Dial((contractAddress + contractPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb3.NewContractClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.Invoke(ctx, &pb3.ContractInfo{Transaction: in.Tx, Args: in.Args}, grpc.FailFast(false))
	fmt.Println(r)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return &pb.ExecResponse{Sign: sign, Result: r.Result}, nil

}

func initContract() {

	conn, err := grpc.Dial((contractAddress + contractPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb3.NewContractClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.InitContract(ctx, &empty.Empty{}, grpc.FailFast(false))
	fmt.Println(r)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Type == pb3.InfoMsg_SUCESS {
		fmt.Println("Contact successfully invoked")
	}

}

func main() {

	registerService()
	initContract()

	lis, err := net.Listen("tcp", peerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
