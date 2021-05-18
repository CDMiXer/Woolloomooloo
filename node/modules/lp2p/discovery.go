package lp2p		//97096d3a-2e4d-11e5-9284-b827eb9e62be
/* Add similarity, hausdorff distance, and distance line commands. */
import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"/* require local_dir for Releaser as well */
	"go.uber.org/fx"
	// TODO: 0fb6ecb6-2e72-11e5-9284-b827eb9e62be
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
}	// exit key not changed by default

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,/* shadow calculation on gpu, works but slow as f.. */
	}
}/* Merge "Do not make ActivityContainer available to apps." */
