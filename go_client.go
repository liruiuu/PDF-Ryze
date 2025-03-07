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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"

	// "log"
	"time"

	pb "pdfguru/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	// name = flag.String("name", defaultName, "Name to greet")
)

func go_client(myStrings []string) {
	flag.Parse()

	for i := 0; i < 15; i++ {
		// Set up a connection to the server.
		conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
		defer cancel()

		log.Printf("i++: %d", i)

		r, err := c.SayHello(ctx, &pb.HelloRequest{MyStrings: myStrings})
		if err == nil {
			log.Printf("Greeting: %s", r.GetMessage())
			break
		} else {
			if myStrings[0] == "closeServer" {
				break
			}
			time.Sleep(2 * time.Second)
			// log.Fatalf("could not greet: %v", err)
		}

	}

}
