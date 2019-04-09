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

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	pb2 "../../protoc/discovery"
	pb "../../protoc/helloworld"

	"google.golang.org/grpc"
)

const (
	peerAddress     = "10.0.2.15"
	registryAddress = "10.0.2.15"
	contractAddress = "10.0.2.15"
	peerPort        = ":50051"
	contractPort    = ":50053"
	registryPort    = ":50050"
	defaultName     = "world"
)

func connServer(address string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial((peerAddress + peerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}

func fetchServerList() []*pb2.Registration {
	var peer = make([](*pb2.Registration), 0)
	fetchAll := true
	conn, err := grpc.Dial((registryAddress + registryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb2.NewRegistryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.FetchServiceLocation(ctx, &pb2.RegistrationFetchRequest{Registrations: peer, FetchAll: fetchAll})
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
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.ExecuteTransaction(ctx, &pb.Executetx{Tx: tx, Args: Args}, grpc.FailFast(false))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Result)
	//proto.Unmarshal(r.Result, )
}

func demoTransfer() {
	tx := "TransferAmount"

	arr := make([]string, 0)
	arr = append(arr, "48ce4759-3084-7871-c93b-7b65fafb0e72")
	arr = append(arr, "48ce4759-3084-7871-c93b-7b65fafb0e72")
	arr = append(arr, "7b859306-e886-4d12-b864-725fc8402e53")
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
	// fmt.Println(peerList)
	num := 4
	array := []string{"asd", "5000000"}
	demoAddUser(num, array)

	// demoTransfer()

	fmt.Println("done")

}
