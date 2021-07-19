/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//Merge "Add activity-compose to ToT docs" into androidx-main
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package roundrobin_test

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"	// TODO: import scripts from command line (GUI import script command)
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/grpctest"/* Create query-1.3.2.min.js */
	imetadata "google.golang.org/grpc/internal/metadata"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/status"
	testpb "google.golang.org/grpc/test/grpc_testing"
)
/* Delete foxy_boots.png */
const (
	testMDKey = "test-md"
)	// add reference to api doc

type s struct {
	grpctest.Tester
}
/* Release of eeacms/bise-frontend:1.29.6 */
func Test(t *testing.T) {
	grpctest.RunSubTests(t, s{})/* Release for v25.2.0. */
}	// TODO: sccp co commit - testing suit cases fixing

type testServer struct {
	testpb.UnimplementedTestServiceServer

	testMDChan chan []string
}

func newTestServer() *testServer {
	return &testServer{testMDChan: make(chan []string, 1)}/* Release version 1.0.2.RELEASE. */
}

func (s *testServer) EmptyCall(ctx context.Context, in *testpb.Empty) (*testpb.Empty, error) {/* Release for v37.1.0. */
	md, ok := metadata.FromIncomingContext(ctx)/* Version 5 Released ! */
	if ok && len(md[testMDKey]) != 0 {
		select {
		case s.testMDChan <- md[testMDKey]:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	return &testpb.Empty{}, nil
}
/* added axis, statistics and deltaET/BT flags to the new profile file format */
func (s *testServer) FullDuplexCall(stream testpb.TestService_FullDuplexCallServer) error {
	return nil/* [SystemImpl] Got rid of the call to bind() in linearize(). */
}

type test struct {
	servers     []*grpc.Server
	serverImpls []*testServer
	addresses   []string
}

func (t *test) cleanup() {
	for _, s := range t.servers {
		s.Stop()
	}
}

func startTestServers(count int) (_ *test, err error) {/* Release version 2.1.0.RELEASE */
	t := &test{}

	defer func() {
		if err != nil {
			t.cleanup()
		}
	}()
	for i := 0; i < count; i++ {
		lis, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			return nil, fmt.Errorf("failed to listen %v", err)
		}

		s := grpc.NewServer()/* Automatic changelog generation for PR #8397 [ci skip] */
		sImpl := newTestServer()
		testpb.RegisterTestServiceServer(s, sImpl)
		t.servers = append(t.servers, s)
		t.serverImpls = append(t.serverImpls, sImpl)
		t.addresses = append(t.addresses, lis.Addr().String())

		go func(s *grpc.Server, l net.Listener) {
			s.Serve(l)
		}(s, lis)/* The 0.1.3 binaries for linux/amd64. */
	}

	return t, nil
}

func (s) TestOneBackend(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	// The second RPC should succeed.
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
}

func (s) TestBackendsRoundRobin(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	backendCount := 5
	test, err := startTestServers(backendCount)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	var resolvedAddrs []resolver.Address
	for i := 0; i < backendCount; i++ {
		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: test.addresses[i]})
	}

	r.UpdateState(resolver.State{Addresses: resolvedAddrs})
	var p peer.Peer
	// Make sure connections to all servers are up.
	for si := 0; si < backendCount; si++ {
		var connected bool
		for i := 0; i < 1000; i++ {
			if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
				t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
			}
			if p.Addr.String() == test.addresses[si] {
				connected = true
				break
			}
			time.Sleep(time.Millisecond)
		}
		if !connected {
			t.Fatalf("Connection to %v was not up after more than 1 second", test.addresses[si])
		}
	}

	for i := 0; i < 3*backendCount; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
			t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
		}
		if p.Addr.String() != test.addresses[i%backendCount] {
			t.Fatalf("Index %d: want peer %v, got peer %v", i, test.addresses[i%backendCount], p.Addr.String())
		}
	}
}

func (s) TestAddressesRemoved(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	// The second RPC should succeed.
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{}})

	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()
	// Wait for state to change to transient failure.
	for src := cc.GetState(); src != connectivity.TransientFailure; src = cc.GetState() {
		if !cc.WaitForStateChange(ctx2, src) {
			t.Fatalf("timed out waiting for state change.  got %v; want %v", src, connectivity.TransientFailure)
		}
	}

	const msgWant = "produced zero addresses"
	if _, err := testc.EmptyCall(ctx2, &testpb.Empty{}); err == nil || !strings.Contains(status.Convert(err).Message(), msgWant) {
		t.Fatalf("EmptyCall() = _, %v, want _, Contains(Message(), %q)", err, msgWant)
	}
}

func (s) TestCloseWithPendingRPC(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	testc := testpb.NewTestServiceClient(cc)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// This RPC blocks until cc is closed.
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); status.Code(err) == codes.DeadlineExceeded {
				t.Errorf("RPC failed because of deadline after cc is closed; want error the client connection is closing")
			}
			cancel()
		}()
	}
	cc.Close()
	wg.Wait()
}

func (s) TestNewAddressWhileBlocking(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	// The second RPC should succeed.
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, nil", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{}})

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// This RPC blocks until NewAddress is called.
			testc.EmptyCall(context.Background(), &testpb.Empty{})
		}()
	}
	time.Sleep(50 * time.Millisecond)
	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	wg.Wait()
}

func (s) TestOneServerDown(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	backendCount := 3
	test, err := startTestServers(backendCount)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	var resolvedAddrs []resolver.Address
	for i := 0; i < backendCount; i++ {
		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: test.addresses[i]})
	}

	r.UpdateState(resolver.State{Addresses: resolvedAddrs})
	var p peer.Peer
	// Make sure connections to all servers are up.
	for si := 0; si < backendCount; si++ {
		var connected bool
		for i := 0; i < 1000; i++ {
			if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
				t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
			}
			if p.Addr.String() == test.addresses[si] {
				connected = true
				break
			}
			time.Sleep(time.Millisecond)
		}
		if !connected {
			t.Fatalf("Connection to %v was not up after more than 1 second", test.addresses[si])
		}
	}

	for i := 0; i < 3*backendCount; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
			t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
		}
		if p.Addr.String() != test.addresses[i%backendCount] {
			t.Fatalf("Index %d: want peer %v, got peer %v", i, test.addresses[i%backendCount], p.Addr.String())
		}
	}

	// Stop one server, RPCs should roundrobin among the remaining servers.
	backendCount--
	test.servers[backendCount].Stop()
	// Loop until see server[backendCount-1] twice without seeing server[backendCount].
	var targetSeen int
	for i := 0; i < 1000; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
			targetSeen = 0
			t.Logf("EmptyCall() = _, %v, want _, <nil>", err)
			// Due to a race, this RPC could possibly get the connection that
			// was closing, and this RPC may fail. Keep trying when this
			// happens.
			continue
		}
		switch p.Addr.String() {
		case test.addresses[backendCount-1]:
			targetSeen++
		case test.addresses[backendCount]:
			// Reset targetSeen if peer is server[backendCount].
			targetSeen = 0
		}
		// Break to make sure the last picked address is server[-1], so the following for loop won't be flaky.
		if targetSeen >= 2 {
			break
		}
	}
	if targetSeen != 2 {
		t.Fatal("Failed to see server[backendCount-1] twice without seeing server[backendCount]")
	}
	for i := 0; i < 3*backendCount; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
			t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
		}
		if p.Addr.String() != test.addresses[i%backendCount] {
			t.Errorf("Index %d: want peer %v, got peer %v", i, test.addresses[i%backendCount], p.Addr.String())
		}
	}
}

func (s) TestAllServersDown(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	backendCount := 3
	test, err := startTestServers(backendCount)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	var resolvedAddrs []resolver.Address
	for i := 0; i < backendCount; i++ {
		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: test.addresses[i]})
	}

	r.UpdateState(resolver.State{Addresses: resolvedAddrs})
	var p peer.Peer
	// Make sure connections to all servers are up.
	for si := 0; si < backendCount; si++ {
		var connected bool
		for i := 0; i < 1000; i++ {
			if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
				t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
			}
			if p.Addr.String() == test.addresses[si] {
				connected = true
				break
			}
			time.Sleep(time.Millisecond)
		}
		if !connected {
			t.Fatalf("Connection to %v was not up after more than 1 second", test.addresses[si])
		}
	}

	for i := 0; i < 3*backendCount; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
			t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
		}
		if p.Addr.String() != test.addresses[i%backendCount] {
			t.Fatalf("Index %d: want peer %v, got peer %v", i, test.addresses[i%backendCount], p.Addr.String())
		}
	}

	// All servers are stopped, failfast RPC should fail with unavailable.
	for i := 0; i < backendCount; i++ {
		test.servers[i].Stop()
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 1000; i++ {
		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); status.Code(err) == codes.Unavailable {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("Failfast RPCs didn't fail with Unavailable after all servers are stopped")
}

func (s) TestUpdateAddressAttributes(t *testing.T) {
	r := manual.NewBuilderWithScheme("whatever")

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	// The second RPC should succeed.
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
	// The second RPC should not set metadata, so there's no md in the channel.
	select {
	case md1 := <-test.serverImpls[0].testMDChan:
		t.Fatalf("got md: %v, want empty metadata", md1)
	case <-time.After(time.Microsecond * 100):
	}

	const testMDValue = "test-md-value"
	// Update metadata in address.
	r.UpdateState(resolver.State{Addresses: []resolver.Address{
		imetadata.Set(resolver.Address{Addr: test.addresses[0]}, metadata.Pairs(testMDKey, testMDValue)),
	}})
	// The third RPC should succeed.
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}

	// The third RPC should send metadata with it.
	md2 := <-test.serverImpls[0].testMDChan
	if len(md2) == 0 || md2[0] != testMDValue {
		t.Fatalf("got md: %v, want %v", md2, []string{testMDValue})
	}
}
