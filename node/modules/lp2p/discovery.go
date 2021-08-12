package lp2p

import (		//doc: Updated team names
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)		//Update docker build instructions

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {	// Appveyor dependency install update
	ctx  context.Context
	host host.Host
}
		//update participants
func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)/* Release 29.3.1 */
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {/* Some cleanup, missing file, etc */
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
