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
//../..
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	// pb3 "study/GitHub/consensus/protoc/contractcode"
	// pb "study/GitHub/consensus/protoc/core"
	// pb2 "study/GitHub/consensus/protoc/discovery"
	// pb5 "study/GitHub/consensus/protoc/gossip"
	// pb4 "study/GitHub/consensus/protoc/orderer"

	pb3 "protoc/contractcode"
	pb "protoc/core"
	pb2 "protoc/discovery"
	pb5 "protoc/gossip"
	pb4 "protoc/orderer"

	"../validate"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	registryAddress  = "10.0.2.15"
	contractAddress  = "10.0.2.15"
	ordererAddress   = "10.0.2.15"
	defaultName      = "10.0.2.15"
	discoveryAddress = "10.0.2.15"
	discoveryPort    = ":50050"
	peerPort         = ":50051"
	contractPort     = ":50053"
	registryPort     = ":50050"
	ordererPort      = ":50054"
	sign             = "Sign"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

var contractClient pb3.ContractClient
var contractConn *grpc.ClientConn

var ordererClient pb4.OrdererClient
var ordererConn *grpc.ClientConn

var discoveryClient pb2.RegistryClient
var discoveryConn *grpc.ClientConn

type peer struct {
	peerClient pb.PeerManagerClient
	peerConn   *grpc.ClientConn
}

var peerList = make(map[string]peer, 0)

func initContract() {

	timeStamp := time.Now()
	conn, err := grpc.Dial((contractAddress + contractPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb3.NewContractClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := c.InitContract(ctx, &empty.Empty{}, grpc.FailFast(false))

	for {
		a, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Unexpected Error", err)
			break
		}
		fmt.Println("Data From Stream :  ", a)
	}
	fmt.Println("Data From Stream :  ", timeStamp)
	// fmt.Println(r)
	if err != nil {
		log.Fatalf("could not stupid greet: %v", err)
	}

}

func (s *server) InitContract(ctx context.Context, in *empty.Empty) (*pb.ExecResponse, error) {

	if contractClient == nil {
		connectContractService()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := contractClient.InitContract(ctx, &empty.Empty{}, grpc.FailFast(false))

	if err != nil {
		log.Fatalf("%v", err)
		return &pb.ExecResponse{Sign: sign, Result: nil}, err
	}

	a := make([]*pb3.TransactionResponse, 0)
	for {
		a, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Unexpected Error", err)
			break
		}
		fmt.Println("Data From Stream :  ", a)
	}

	res := &pb3.Cargo{Load: a}
	cargo, err := proto.Marshal(res)
	arr := make([]string, 3)
	submitTx(arr, cargo)
	return &pb.ExecResponse{Sign: sign, Result: nil}, err
}

func (s *server) ExecuteTransaction(ctx context.Context, in *pb.Executetx) (*pb.ExecResponse, error) {

	var c chan *pb.ExecResponse
	if in.IType == pb.Executetx_CLIENT {
		c = make(chan *pb.ExecResponse)
		go endorceTx(in, c)
	}

	if contractClient == nil {
		connectContractService()
	}
	// conn, err := grpc.Dial((contractAddress + contractPort), grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// c := pb3.NewContractClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stream, err := contractClient.Invoke(ctx, &pb3.ContractInfo{Transaction: in.Tx, Args: in.Args}, grpc.FailFast(false))
	//fmt.Println(r)
	if err != nil {
		log.Fatalf("%v", err)
		return &pb.ExecResponse{Sign: sign, Result: nil}, err
	}

	a := make([]*pb3.TransactionResponse, 0)
	for {
		tmp, err := stream.Recv()
		a = append(a, tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Unexpected Error", err)
			break
		}
		fmt.Println("Data From Stream :  ", a)
	}
	// // fmt.Println(r)
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	res := &pb3.Cargo{Load: a}
	cargo, err := proto.Marshal(res)

	if in.IType == pb.Executetx_CLIENT {
		signs := make([]string, 3)
		// for i := 0; i < 3; i++ {
		// 	temp := <-c
		// 	signs[i] = temp.Sign
		// }

		submitTx(signs, cargo)
		fmt.Println("tx submitted")
		return &pb.ExecResponse{Sign: sign, Result: nil}, err
	}
	fmt.Println("tx not submitted")
	return &pb.ExecResponse{Sign: sign, Result: cargo}, err

}

func (s *server) ValidateBlock(ctx context.Context, in *pb5.Block) (*empty.Empty, error) {

	blk := new(validate.BlockInfo)

	blk.SetEndoresedTxs(in.Bunch)
	blk.SetSign(in.Sign)
	blk.SetOffset(in.OffSet)
	blk.SetMask(len(in.Bunch))
	blk.ValidateBlock()
	fmt.Println("block validated")
	return &empty.Empty{}, nil

}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.Name)
// 	time.Sleep(2 * time.Second)
// 	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
// }

// func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
// }

func submitTx(signs []string, cargo []byte) {
	conn, err := grpc.Dial((ordererAddress + ordererPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb4.NewOrdererClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// t := time.Now().String
	z, err := c.SubmitTx(ctx, &pb4.EndorsedTx{Sign: signs, Payload: cargo}, grpc.FailFast(false))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		fmt.Println(z)
	}

}

func endorceTx(in *pb.Executetx, c chan *pb.ExecResponse) {
	// var err error
	registerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if discoveryClient == nil {
		createDiscoveryConnection()
	}
	z, err := discoveryClient.FetchServiceLocation(registerCtx, &empty.Empty{}, grpc.FailFast(false))

	if err != nil {
		return
	}
	for i := 0; i < len(z.Registrations); i++ {
		if z.Registrations[i].Ipv4 != defaultName {
			go invokePeer(z.Registrations[i], in, c)
		}
	}
}

func invokePeer(z *pb2.Registration, in *pb.Executetx, c chan *pb.ExecResponse) {
	if _, ok := peerList[(z.Ipv4)]; ok {
		peerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		tempClient := peerList[(z.Ipv4)]
		response, err := tempClient.peerClient.ExecuteTransaction(peerCtx, &pb.Executetx{Tx: in.Tx, Args: in.Args, InvokerId: in.InvokerId, IType: pb.Executetx_PEER}, grpc.FailFast(false))
		if err != nil {
			cancel()
			return
		}
		c <- response
		cancel()

	} else {
		connectPeer(z)
		peerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		tempClient := peerList[(z.Ipv4)]
		response, err := tempClient.peerClient.ExecuteTransaction(peerCtx, &pb.Executetx{Tx: in.Tx, Args: in.Args, InvokerId: in.InvokerId, IType: pb.Executetx_PEER}, grpc.FailFast(false))
		if err != nil {
			cancel()
			return
		}
		c <- response
		cancel()

	}

}

// func peerInvoke(peerAddress string, peerPort string) {
// 	conn, err := grpc.Dial((peerAddress + peerPort), grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewPeerManagerClient(conn)

// 	// Contact the server and print out its response.

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	r, err := c.ExecuteTransaction(ctx, &pb.Executetx{Tx: tx, Args: Args}, grpc.FailFast(false))
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.Result)
// }

func connectPeer(peerInfo *pb2.Registration) {
	var err error
	tempConn, err := grpc.Dial((peerInfo.Ipv4 + peerInfo.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	tempClient := pb.NewPeerManagerClient(tempConn)
	peerList[peerInfo.Ipv4] = peer{
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

func connectOrdererService() {
	var err error
	ordererConn, err = grpc.Dial((ordererAddress + ordererPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ordererClient = pb4.NewOrdererClient(ordererConn)
}

func closeOrdererService() {
	ordererConn.Close()
	ordererConn = nil
	ordererClient = nil
}

func connectContractService() {
	var err error
	contractConn, err = grpc.Dial((contractAddress + contractPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	contractClient = pb3.NewContractClient(contractConn)
}

func closeContractService() {
	contractConn.Close()
	contractConn = nil
	contractClient = nil
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

func main() {

	registerService()
	initContract()

	lis, err := net.Listen("tcp", peerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPeerManagerServer(s, &server{})
	fmt.Println("registered")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// connectOrdererService()
	// connectContractService()
}
