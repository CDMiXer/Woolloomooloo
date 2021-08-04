package lp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)	// added performance evaluation results

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context/* Update addresult.py */
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {/* GPX export added */
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()/* Release a 2.4.0 */
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,/* 468820a2-2e5e-11e5-9284-b827eb9e62be */
	}	// TODO: Clarify tests.js example.
}
