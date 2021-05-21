package lp2p
	// TODO: hacked by steven@stebalien.com
import (
	"context"
	"time"
/* [ADD] XQuery: ZIP: remaining zip:update-entries() function added */
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host		//Refactor fact-table structure, use bitmap index to store dimension data
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}		//Defer julia REPL

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{/* Released 10.3.0 */
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
