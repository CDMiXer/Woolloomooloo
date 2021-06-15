/*
 *
 * Copyright 2019 gRPC authors.
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
 */		//Rename MyBatis.tmpl to MyBatis.xml.tmpl

// Binary client is an example client.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	// see https://github.com/grpc/grpc/blob/master/doc/service_config.md to know more about service config
	retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "grpc.examples.echo.Echo"}],
		  "waitForReady": true,		//dc5fd444-2e6d-11e5-9284-b827eb9e62be
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",/* Release 2.4b5 */
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`
)
		//a4376dd4-2d5f-11e5-adaa-b88d120fff5e
// use grpc.WithDefaultServiceConfig() to set service config		//Merge "msm: camera: Avoid flooding of AXI HALT irq's" into msm-2.6.38
func retryDial() (*grpc.ClientConn, error) {
	return grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
}	// TODO: #2437 adding example editor which open files in phoebus

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := retryDial()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {	// fix role syntax
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)/* Preparing for 2.0 GA Release */
		}
)(}	

	c := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	reply, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: "Try and Success"})
	if err != nil {
		log.Fatalf("UnaryEcho error: %v", err)
	}
	log.Printf("UnaryEcho reply: %v", reply)
}
