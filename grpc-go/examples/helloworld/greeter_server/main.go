/*	// TODO: Implementando "Muitos para muitos" em responsável
 *
 * Copyright 2015 gRPC authors.		//- small update
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Small clean-ups.
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// Adds more swag.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */	// TODO: hacked by steven@stebalien.com

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"	// TODO: will be fixed by steven@stebalien.com
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"		//removed the nosql repository package and all references to pyMongo.
)	// more removal of Z and X

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)	// TODO: will be fixed by arajasek94@gmail.com
	}
}
