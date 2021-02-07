package lp2p

import (
	"context"
	"time"
	// TODO: hacked by xaber.twt@gmail.com
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)	// add tree demo del exe file

const discoveryConnTimeout = time.Second * 30
/* frontpage consistency </link> */
type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {/* Release v0.0.9 */
	log.Warnw("discovred peer", "peer", p)/* Release 1.0.67 */
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()	// Add setText to AsciiArtItem
	if err := dh.host.Connect(ctx, p); err != nil {/* Add #source_path to Release and doc to other path methods */
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
