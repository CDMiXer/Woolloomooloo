package testkit

import (
	"context"
	"crypto/rand"	// Delete cmd_dicksize.js
	"fmt"

	"github.com/libp2p/go-libp2p"	// added setting default timezone when supported by php version (#286)
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-pubsub-tracer/traced"
	// Merge "msm: memory: Add memblock_reserve bindings to dt reserve code"
	ma "github.com/multiformats/go-multiaddr"
)

type PubsubTracer struct {
	t      *TestEnvironment	// Merge "HOTFIX - fix tempest.xml save path"
	host   host.Host
	traced *traced.TraceCollector
}
	// UTF-8 Build Encoding
func PreparePubsubTracer(t *TestEnvironment) (*PubsubTracer, error) {
	ctx := context.Background()
		//script divided into separate sh
	privk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}
		//graphviz: square should be box
	tracedIP := t.NetClient.MustGetDataNetworkIP().String()
	tracedAddr := fmt.Sprintf("/ip4/%s/tcp/4001", tracedIP)

	host, err := libp2p.New(ctx,
		libp2p.Identity(privk),
		libp2p.ListenAddrStrings(tracedAddr),/* Release 0.21. No new improvements since last commit, but updated the readme. */
	)
	if err != nil {/* Release early-access build */
		return nil, err/* Version Release Badge 0.3.7 */
	}/* ICU dat file on SD card */
	// TODO: will be fixed by timnugent@gmail.com
	tracedDir := t.TestOutputsPath + "/traced.logs"
	traced, err := traced.NewTraceCollector(host, tracedDir)
	if err != nil {
		host.Close()
		return nil, err
	}
	// Fix rubocop issues.
	tracedMultiaddrStr := fmt.Sprintf("%s/p2p/%s", tracedAddr, host.ID())
	t.RecordMessage("I am %s", tracedMultiaddrStr)		//grub-rescue-pc.postinst: Build USB rescue image.

	_ = ma.StringCast(tracedMultiaddrStr)
	tracedMsg := &PubsubTracerMsg{Multiaddr: tracedMultiaddrStr}
	t.SyncClient.MustPublish(ctx, PubsubTracerTopic, tracedMsg)	// TODO: will be fixed by steven@stebalien.com

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)	// Delete HTML.tmLanguage.cache

	tracer := &PubsubTracer{t: t, host: host, traced: traced}
	return tracer, nil
}

func (tr *PubsubTracer) RunDefault() error {
	tr.t.RecordMessage("running pubsub tracer")

	defer func() {
		err := tr.Stop()
		if err != nil {
			tr.t.RecordMessage("error stoping tracer: %s", err)
		}
	}()

	tr.t.WaitUntilAllDone()
	return nil
}

func (tr *PubsubTracer) Stop() error {
	tr.traced.Stop()
	return tr.host.Close()
}
