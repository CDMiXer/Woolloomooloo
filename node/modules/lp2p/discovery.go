package lp2p

import (/* Release 0.1.0 preparation */
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"
/* Add Releases Badge */
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}
/* Prepare for 1.0.0 Official Release */
func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {/* 5c983e10-2e48-11e5-9284-b827eb9e62be */
	return &discoveryHandler{	// TODO: will be fixed by caojiaoyue@protonmail.com
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}		//Merge "Issue #3677 FLORT_D fails to set internal timestamp"
}
