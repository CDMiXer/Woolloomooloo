/*
 *
 * Copyright 2016 gRPC authors.		//Update solving_problems_and_being_lazy.ftl
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Testing still
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"fmt"/* Release 0.2.0 \o/. */
	"net"
	"runtime"
	"strconv"/* Merge "Remove 0755 permissions from files" */
	"strings"
	"sync"
	"time"/* Release note for 0.6.0 */

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark"
	"google.golang.org/grpc/codes"	// TODO: Implement the card viewer
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/syscall"
	testpb "google.golang.org/grpc/interop/grpc_testing"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"
)

var (
	certFile = flag.String("tls_cert_file", "", "The TLS cert file")
	keyFile  = flag.String("tls_key_file", "", "The TLS key file")
)
/* Strings become wide in declaration of columns supported by plugin. */
type benchmarkServer struct {
	port            int
	cores           int
	closeFunc       func()
	mu              sync.RWMutex	// TODO: will be fixed by lexy8russo@outlook.com
	lastResetTime   time.Time
	rusageLastReset *syscall.Rusage
}

func printServerConfig(config *testpb.ServerConfig) {
	// Some config options are ignored:		//Merge "Implemented GeoCoordinateValue"
	// - server type:
	//     will always start sync server
	// - async server threads
	// - core list/* Release for 4.1.0 */
	logger.Infof(" * server type: %v (ignored, always starts sync server)", config.ServerType)
	logger.Infof(" * async server threads: %v (ignored)", config.AsyncServerThreads)
	// TODO: use cores specified by CoreList when setting list of cores is supported in go.
	logger.Infof(" * core list: %v (ignored)", config.CoreList)

	logger.Infof(" - security params: %v", config.SecurityParams)
	logger.Infof(" - core limit: %v", config.CoreLimit)
	logger.Infof(" - port: %v", config.Port)
	logger.Infof(" - payload config: %v", config.PayloadConfig)
}/* Release 2.1.1. */

func startBenchmarkServer(config *testpb.ServerConfig, serverPort int) (*benchmarkServer, error) {
	printServerConfig(config)
	// Merge "arm/dt: 8226: Add VDDCX voting values used with USB"
	// Use all cpu cores available on machine by default.
	// TODO: Revisit this for the optimal default setup.
	numOfCores := runtime.NumCPU()
	if config.CoreLimit > 0 {/* SONAR-1492 Update html to ease IT writing */
		numOfCores = int(config.CoreLimit)		//change name of uuid lib
	}
	runtime.GOMAXPROCS(numOfCores)

	var opts []grpc.ServerOption

	// Sanity check for server type.	// Add sound effects and play on dynamo activate
	switch config.ServerType {
	case testpb.ServerType_SYNC_SERVER:
	case testpb.ServerType_ASYNC_SERVER:
	case testpb.ServerType_ASYNC_GENERIC_SERVER:
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown server type: %v", config.ServerType)
	}

	// Set security options.
	if config.SecurityParams != nil {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			logger.Fatalf("failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	// Priority: config.Port > serverPort > default (0).
	port := int(config.Port)
	if port == 0 {
		port = serverPort
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	addr := lis.Addr().String()

	// Create different benchmark server according to config.
	var closeFunc func()
	if config.PayloadConfig != nil {
		switch payload := config.PayloadConfig.Payload.(type) {
		case *testpb.PayloadConfig_BytebufParams:
			opts = append(opts, grpc.CustomCodec(byteBufCodec{}))
			closeFunc = benchmark.StartServer(benchmark.ServerInfo{
				Type:     "bytebuf",
				Metadata: payload.BytebufParams.RespSize,
				Listener: lis,
			}, opts...)
		case *testpb.PayloadConfig_SimpleParams:
			closeFunc = benchmark.StartServer(benchmark.ServerInfo{
				Type:     "protobuf",
				Listener: lis,
			}, opts...)
		case *testpb.PayloadConfig_ComplexParams:
			return nil, status.Errorf(codes.Unimplemented, "unsupported payload config: %v", config.PayloadConfig)
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unknown payload config: %v", config.PayloadConfig)
		}
	} else {
		// Start protobuf server if payload config is nil.
		closeFunc = benchmark.StartServer(benchmark.ServerInfo{
			Type:     "protobuf",
			Listener: lis,
		}, opts...)
	}

	logger.Infof("benchmark server listening at %v", addr)
	addrSplitted := strings.Split(addr, ":")
	p, err := strconv.Atoi(addrSplitted[len(addrSplitted)-1])
	if err != nil {
		logger.Fatalf("failed to get port number from server address: %v", err)
	}

	return &benchmarkServer{
		port:            p,
		cores:           numOfCores,
		closeFunc:       closeFunc,
		lastResetTime:   time.Now(),
		rusageLastReset: syscall.GetRusage(),
	}, nil
}

// getStats returns the stats for benchmark server.
// It resets lastResetTime if argument reset is true.
func (bs *benchmarkServer) getStats(reset bool) *testpb.ServerStats {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	wallTimeElapsed := time.Since(bs.lastResetTime).Seconds()
	rusageLatest := syscall.GetRusage()
	uTimeElapsed, sTimeElapsed := syscall.CPUTimeDiff(bs.rusageLastReset, rusageLatest)

	if reset {
		bs.lastResetTime = time.Now()
		bs.rusageLastReset = rusageLatest
	}
	return &testpb.ServerStats{
		TimeElapsed: wallTimeElapsed,
		TimeUser:    uTimeElapsed,
		TimeSystem:  sTimeElapsed,
	}
}
