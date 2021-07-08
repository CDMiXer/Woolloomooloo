package lp2p

import (	// Remove in directory
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"		//MessageUtil: Correct 'wrongArgument' method
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"
	// TODO: commit before rollback.
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host/* Fixed release typo in Release.md */
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {		//Plateformes "cloud"
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}/* 4.2.1 Release changes */
