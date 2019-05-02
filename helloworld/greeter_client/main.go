package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "study/GitHub/consensus/protoc/core"
	pb2 "study/GitHub/consensus/protoc/discovery"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

const (
	peerAddress      = "10.20.24.26"
	registryAddress  = "10.20.24.26"
	contractAddress  = "10.20.24.26"
	discoveryAddress = "10.20.24.26"
	discoveryPort    = ":50050"
	peerPort         = ":50051"
	contractPort     = ":50053"
	registryPort     = ":50050"
	defaultName      = "world"
)

func fetchServerList() []*pb2.Registration {
	// var peer = make([](*pb2.Registration), 0)
	// fetchAll := true
	conn, err := grpc.Dial((registryAddress + registryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb2.NewRegistryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.FetchServiceLocation(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println(r)

	return r.Registrations
}

func invokeTransaction(tx string, Args []string) {

	conn, err := grpc.Dial((peerAddress + peerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPeerManagerClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.ExecuteTransaction(ctx, &pb.Executetx{Tx: tx, Args: Args, InvokerId: "userID", IType: pb.Executetx_CLIENT}, grpc.FailFast(false))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Result)
	//proto.Unmarshal(r.Result, )
}

func demoTransfer() {
	tx := "TransferAmount"

	arr := make([]string, 0)
	arr = append(arr, "42915f0f-11eb-431d-9416-4ea2da3b85bb")
	arr = append(arr, "42915f0f-11eb-431d-9416-4ea2da3b85bb")
	arr = append(arr, "6a21d603-8fda-48b7-b7f1-bee772979a40")
	arr = append(arr, "500")
	// a.FromAccId = "790efe70-80f9-68c5-696d-c23a6552868c"
	// a.ToAccId = "7c3aaa44-0ab7-8abe-35db-871c376e968f"
	// a.UserId = "790efe70-80f9-68c5-696d-c23a6552868c"

	invokeTransaction(tx, arr)
}

func demoAddUser(num int, arr []string) {

	tx := "AddUser"
	// var a pb3.UserInfo

	// for i := 1; i <= num; i++ {
	// 	a.UserName = arr[i-1]
	// 	a.Balance = 5000
	// 	//Args, err := proto.Marshal(&a)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	invokeTransaction(tx, arr)
	// }
}

func main() {
	// var peerList = make([](*pb2.Registration), 0)
	// peerList = fetchServerList()
	// 	fmt.Println(peerList)
	// num := 4
	// array := []string{"asd", "5000000"}
	// demoAddUser(num, array)

	demoTransfer()

	fmt.Println("done")

}
