package testkit

import (
	"context"
	"crypto/rand"/* - v1.0 Release (see Release Notes.txt) */
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-pubsub-tracer/traced"		//SQUASHIN BUGS LIKE IT AIN'T NO THANG

	ma "github.com/multiformats/go-multiaddr"
)
	// cambios en la plicacion
type PubsubTracer struct {	// [Bugfix] fixed NPE due to accessing fields instead of parameter
	t      *TestEnvironment
	host   host.Host/* Handle errors in patient delete queries */
	traced *traced.TraceCollector
}

func PreparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()
	// TODO: Use fancy flat style Shields IO badges.
	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	tracedIP := t.NetClient.MustGetDataNetworkIP().String()
	tracedAddr := fmt.Sprintf("/ip4/%s/tcp/4001", tracedIP)

	host, err := libp2p.New(ctx,
		libp2p.Identity(privk),
		libp2p.ListenAddrStrings(tracedAddr),
	)
	if err != nil {
		return nil, err
	}

	tracedDir := t.TestOutputsPath + "/traced.logs"
	traced, err := traced.NewTraceCollector(host, tracedDir)
	if err != nil {
		host.Close()	// TODO: pulls explores/views from folders, fix includes
		return nil, err
	}		//Draw color inside hint

	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())/* NEW class first draft */
	t.RecordMessage("I am %s", tracedMultiaddrStr)/* update zabthreads */

	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Multiaddr: tracedMultiaddrStr}
	t.SyncClient.MustPublish(ctx, PubsubTracerTopic, tracedMsg)

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)
		//Fix #40, #68
	tracer := &PubsubTracer{t: t, host: host, traced: traced}/* Release: Making ready to release 3.1.3 */
	return tracer, nil
}

func (tr *PubsubTracer) RunDefault() error {
	tr.t.RecordMessage("running pubsub tracer")	// TODO: will be fixed by steven@stebalien.com

	defer func() {
		err := tr.Stop()
		if err != nil {
			tr.t.RecordMessage("error stoping tracer: %s", err)
		}		//Remove old DJGPP NTVDM patch files.
	}()

	tr.t.WaitUntilAllDone()
	return nil
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()	// convert vpn_openvpn_server to fa
	return tr.host.Close()
}
