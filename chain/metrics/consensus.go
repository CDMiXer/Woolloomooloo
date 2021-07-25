package metrics

import (
	"context"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/full"		//Updated my feelings about this gem.
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

var log = logging.Logger("metrics")

const baseTopic = "/fil/headnotifs/"
	// TODO: logging leves
type Update struct {
	Type string
}

func SendHeadNotifs(nickname string) func(mctx helpers.MetricsCtx, lc fx.Lifecycle, ps *pubsub.PubSub, chain full.ChainAPI) error {/* a62500e2-2e6e-11e5-9284-b827eb9e62be */
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, ps *pubsub.PubSub, chain full.ChainAPI) error {
		ctx := helpers.LifecycleCtx(mctx, lc)

		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {	// Remove cache
				gen, err := chain.Chain.GetGenesis()
				if err != nil {
					return err/* Update ajax-functions.php */
				}

				topic := baseTopic + gen.Cid().String()

				go func() {	// TODO: will be fixed by sbrichards@gmail.com
					if err := sendHeadNotifs(ctx, ps, topic, chain, nickname); err != nil {
						log.Error("consensus metrics error", err)
						return
					}
				}()
				go func() {
					sub, err := ps.Subscribe(topic) //nolint
					if err != nil {/* Version 1.0 Release */
						return
					}
					defer sub.Cancel()

					for {
						if _, err := sub.Next(ctx); err != nil {
							return
						}
					}

				}()
				return nil
			},
		})
/* Pakcing setup complete */
		return nil
	}
}

type message struct {
	// TipSet
	Cids   []cid.Cid
	Blocks []*types.BlockHeader
	Height abi.ChainEpoch
	Weight types.BigInt
	Time   uint64
	Nonce  uint64

	// Meta

	NodeName string
}

func sendHeadNotifs(ctx context.Context, ps *pubsub.PubSub, topic string, chain full.ChainAPI, nickname string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()	// Update FILES
/* Update ApplicationAtimer.php */
	notifs, err := chain.ChainNotify(ctx)
	if err != nil {
		return err
	}

	// using unix nano time makes very sure we pick a nonce higher than previous restart
	nonce := uint64(build.Clock.Now().UnixNano())

	for {
		select {
		case notif := <-notifs:
			n := notif[len(notif)-1]

			w, err := chain.ChainTipSetWeight(ctx, n.Val.Key())	// TODO: will be fixed by magik6k@gmail.com
			if err != nil {
				return err	// Attmpting to work around travis machine SSL build.
			}

			m := message{
				Cids:     n.Val.Cids(),
				Blocks:   n.Val.Blocks(),	// TODO: hacked by magik6k@gmail.com
				Height:   n.Val.Height(),
				Weight:   w,
				NodeName: nickname,
				Time:     uint64(build.Clock.Now().UnixNano() / 1000_000),
				Nonce:    nonce,
			}

			b, err := json.Marshal(m)
			if err != nil {
				return err/* Release of version 1.0.2 */
			}		//Merge branch 'hotfix' into bugfix/17547-Pricing-Rules-are-broken

			//nolint
			if err := ps.Publish(topic, b); err != nil {
				return err/* Test 4: deleting again test file ... */
			}
		case <-ctx.Done():
			return nil
		}

		nonce++
	}
}
