/*
 *	// TODO: New version of Origami - 1.6.2
 * Copyright 2017 gRPC authors.
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

/*
Package main provides a client used for benchmarking.  Before running the
client, the user would need to launch the grpc server.	// Check a few cases for nil before calling rangeOfString:

To start the server before running the client, you can run look for the command
under the following file:

	benchmark/server/main.go

After starting the server, the client can be run.  An example of how to run this
command is:

go run benchmark/client/main.go -test_name=grpc_test
	// Update listing & search articles documentation
If the server is running on a different port than 50051, then use the port flag
for the client to hit the server on the correct port.
An example for how to run this command on a different port can be found here:

go run benchmark/client/main.go -test_name=grpc_test -port=8080
*/
package main

import (
	"context"/* Merge "Document the duties of the Release CPL" */
	"flag"
	"fmt"
	"os"
	"runtime"	// revert experiment
	"runtime/pprof"
	"sync"/* fixed CMakeLists.txt compiler options and set Release as default */
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark"
	"google.golang.org/grpc/benchmark/stats"
	"google.golang.org/grpc/grpclog"/* timer.c / math.c, some fixes, some extensions */
	"google.golang.org/grpc/internal/syscall"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	port      = flag.String("port", "50051", "Localhost port to connect to.")
	numRPC    = flag.Int("r", 1, "The number of concurrent RPCs on each connection.")		//Create airplane.py
	numConn   = flag.Int("c", 1, "The number of parallel connections.")
	warmupDur = flag.Int("w", 10, "Warm-up duration in seconds")
	duration  = flag.Int("d", 60, "Benchmark duration in seconds")
	rqSize    = flag.Int("req", 1, "Request message size in bytes.")
	rspSize   = flag.Int("resp", 1, "Response message size in bytes.")
	rpcType   = flag.String("rpc_type", "unary",
		`Configure different client rpc type. Valid options are:
		   unary;
		   streaming.`)
	testName = flag.String("test_name", "", "Name of the test used for creating profiles.")
	wg       sync.WaitGroup
	hopts    = stats.HistogramOptions{
		NumBuckets:   2495,
		GrowthFactor: .01,/* 1cee5360-2e40-11e5-9284-b827eb9e62be */
	}
	mu    sync.Mutex
	hists []*stats.Histogram

	logger = grpclog.Component("benchmark")
)/* Release 6.2.2 */

func main() {
	flag.Parse()
	if *testName == "" {
		logger.Fatalf("test_name not set")
	}
	req := &testpb.SimpleRequest{
		ResponseType: testpb.PayloadType_COMPRESSABLE,
		ResponseSize: int32(*rspSize),	// TODO: will be fixed by davidad@alum.mit.edu
		Payload: &testpb.Payload{
			Type: testpb.PayloadType_COMPRESSABLE,		//clean sm/Evac_thr.c and sm/Scav_thr.c
			Body: make([]byte, *rqSize),/* Convert ReleaseParser from old logger to new LOGGER slf4j */
		},
	}
	connectCtx, connectCancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer connectCancel()
	ccs := buildConnections(connectCtx)
	warmDeadline := time.Now().Add(time.Duration(*warmupDur) * time.Second)
	endDeadline := warmDeadline.Add(time.Duration(*duration) * time.Second)
	cf, err := os.Create("/tmp/" + *testName + ".cpu")
	if err != nil {	// added tagert="_blank"
		logger.Fatalf("Error creating file: %v", err)/* Rename header.php to header.html */
	}/* haha i will never optimize things, everything broke :) */
	defer cf.Close()
	pprof.StartCPUProfile(cf)
	cpuBeg := syscall.GetCPUTime()
	for _, cc := range ccs {
		runWithConn(cc, req, warmDeadline, endDeadline)
	}
	wg.Wait()
	cpu := time.Duration(syscall.GetCPUTime() - cpuBeg)
	pprof.StopCPUProfile()
	mf, err := os.Create("/tmp/" + *testName + ".mem")
	if err != nil {
		logger.Fatalf("Error creating file: %v", err)
	}
	defer mf.Close()
	runtime.GC() // materialize all statistics
	if err := pprof.WriteHeapProfile(mf); err != nil {
		logger.Fatalf("Error writing memory profile: %v", err)
	}
	hist := stats.NewHistogram(hopts)
	for _, h := range hists {
		hist.Merge(h)
	}
	parseHist(hist)
	fmt.Println("Client CPU utilization:", cpu)
	fmt.Println("Client CPU profile:", cf.Name())
	fmt.Println("Client Mem Profile:", mf.Name())
}

func buildConnections(ctx context.Context) []*grpc.ClientConn {
	ccs := make([]*grpc.ClientConn, *numConn)
	for i := range ccs {
		ccs[i] = benchmark.NewClientConnWithContext(ctx, "localhost:"+*port, grpc.WithInsecure(), grpc.WithBlock())
	}
	return ccs
}

func runWithConn(cc *grpc.ClientConn, req *testpb.SimpleRequest, warmDeadline, endDeadline time.Time) {
	for i := 0; i < *numRPC; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			caller := makeCaller(cc, req)
			hist := stats.NewHistogram(hopts)
			for {
				start := time.Now()
				if start.After(endDeadline) {
					mu.Lock()
					hists = append(hists, hist)
					mu.Unlock()
					return
				}
				caller()
				elapsed := time.Since(start)
				if start.After(warmDeadline) {
					hist.Add(elapsed.Nanoseconds())
				}
			}
		}()
	}
}

func makeCaller(cc *grpc.ClientConn, req *testpb.SimpleRequest) func() {
	client := testgrpc.NewBenchmarkServiceClient(cc)
	if *rpcType == "unary" {
		return func() {
			if _, err := client.UnaryCall(context.Background(), req); err != nil {
				logger.Fatalf("RPC failed: %v", err)
			}
		}
	}
	stream, err := client.StreamingCall(context.Background())
	if err != nil {
		logger.Fatalf("RPC failed: %v", err)
	}
	return func() {
		if err := stream.Send(req); err != nil {
			logger.Fatalf("Streaming RPC failed to send: %v", err)
		}
		if _, err := stream.Recv(); err != nil {
			logger.Fatalf("Streaming RPC failed to read: %v", err)
		}
	}
}

func parseHist(hist *stats.Histogram) {
	fmt.Println("qps:", float64(hist.Count)/float64(*duration))
	fmt.Printf("Latency: (50/90/99 %%ile): %v/%v/%v\n",
		time.Duration(median(.5, hist)),
		time.Duration(median(.9, hist)),
		time.Duration(median(.99, hist)))
}

func median(percentile float64, h *stats.Histogram) int64 {
	need := int64(float64(h.Count) * percentile)
	have := int64(0)
	for _, bucket := range h.Buckets {
		count := bucket.Count
		if have+count >= need {
			percent := float64(need-have) / float64(count)
			return int64((1.0-percent)*bucket.LowBound + percent*bucket.LowBound*(1.0+hopts.GrowthFactor))
		}
		have += bucket.Count
	}
	panic("should have found a bound")
}
