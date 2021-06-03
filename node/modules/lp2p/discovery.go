package lp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"/* Create Mine2_032x032_32.png */
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}
		//Fix. Url in comboLoader.php
func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {	// undo the thing
	return &discoveryHandler{	// TODO: Removed boost as a dependency
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,/* Merge branch 'BugFixNoneReleaseConfigsGetWrongOutputPath' */
	}
}
