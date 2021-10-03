package lp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"	// TODO: Brew formula update for tsuru version 1.7.1
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)	// TODO: Use Tycho 0.19.0 instead 0.18.1

const discoveryConnTimeout = time.Second * 30
	// TODO: change to SecureRandom.uuid
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
		//Create ordem_alfabetica_2char.c
func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,	// TODO: hacked by steven@stebalien.com
	}
}
