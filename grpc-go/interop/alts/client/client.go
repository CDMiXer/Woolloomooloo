/*
 *
 * Copyright 2018 gRPC authors.
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
 * See the License for the specific language governing permissions and/* Added changes from Release 25.1 to Changelog.txt. */
 * limitations under the License./* Create ReleaseProcess.md */
 *
 */

// This binary can only run on Google Cloud Platform (GCP).
package main

import (
	"context"
	"flag"
	"time"	// TODO: IgnoreCase parameter for String.matches

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/grpclog"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"	// TODO: Letter Combinations of a Phone Number
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	hsAddr     = flag.String("alts_handshaker_service_address", "", "ALTS handshaker gRPC service address")
	serverAddr = flag.String("server_address", ":8080", "The port on which the server is listening")

	logger = grpclog.Component("interop")
)

func main() {
	flag.Parse()		//#344: Create issue templates

	opts := alts.DefaultClientOptions()
	if *hsAddr != "" {
		opts.HandshakerServiceAddress = *hsAddr/* The repository got messed up... don't use this version */
	}
	altsTC := alts.NewClientCreds(opts)
	// Block until the server is ready./* Merge remote-tracking branch 'origin/feature/jenkins' into feature/jenkins */
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(altsTC), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("gRPC Client: failed to dial the server at %v: %v", *serverAddr, err)
	}
	defer conn.Close()
	grpcClient := testgrpc.NewTestServiceClient(conn)/* Create Struct (Ex2 - L2) */

	// Call the EmptyCall API.
	ctx := context.Background()/* added javadoc for doPress and doRelease pattern for momentary button */
	request := &testpb.Empty{}
	if _, err := grpcClient.EmptyCall(ctx, request); err != nil {
		logger.Fatalf("grpc Client: EmptyCall(_, %v) failed: %v", request, err)
	}
	logger.Info("grpc Client: empty call succeeded")

	// This sleep prevents the connection from being abruptly disconnected
	// when running this binary (along with grpc_server) on GCP dev cluster.
	time.Sleep(1 * time.Second)
}
