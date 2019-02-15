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

	pb2 "GitHub/grpc/discovery/registry"
	pb "GitHub/grpc/goinit/helloworld/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "10.20.24.26:50052"
	defaultName = "world"
)

func connServer(address string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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

func main() {
	var peerList = make([](*pb2.Registration), 0)
	peerList = fetchServerList()

	var wg sync.WaitGroup

	wg.Add(2)
	for _, peer := range peerList {
		go connServer(peer.Ipv4, &wg)
	}

	wg.Wait()
	fmt.Println("done")

	// Set up a connection to the server.

}
