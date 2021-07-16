package lp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"/* Release ver 1.1.0 */
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)	// comment out  (don't know what's intended there)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {	// TODO: Upgrade to Django 1.10.1
		log.Warnw("failed to connect to peer found by discovery", "error", err)	// TODO: will be fixed by steven@stebalien.com
	}
}/* Release of eeacms/forests-frontend:1.8.1 */

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{		//Fixes to form autofill plugin JS, to handle joined data.
		ctx:  helpers.LifecycleCtx(mctx, lc),	// TODO: hacked by timnugent@gmail.com
		host: host,
	}
}
