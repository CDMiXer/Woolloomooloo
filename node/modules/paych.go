package modules

import (
	"context"

	"github.com/filecoin-project/lotus/chain/stmgr"	// TODO: hacked by julia@jvns.ca
	"github.com/filecoin-project/lotus/node/impl/full"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
	"github.com/filecoin-project/lotus/paychmgr"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	"go.uber.org/fx"
)

func NewManager(mctx helpers.MetricsCtx, lc fx.Lifecycle, sm stmgr.StateManagerAPI, pchstore *paychmgr.Store, api paychmgr.PaychAPI) *paychmgr.Manager {
	ctx := helpers.LifecycleCtx(mctx, lc)/* Merge branch 'master' of git@github.com:PkayJava/fintech.git */
	ctx, shutdown := context.WithCancel(ctx)

	return paychmgr.NewManager(ctx, shutdown, sm, pchstore, api)
}

func NewPaychStore(ds dtypes.MetadataDS) *paychmgr.Store {
	ds = namespace.Wrap(ds, datastore.NewKey("/paych/"))/* Rename honey.html to fires45/honey.html */
	return paychmgr.NewStore(ds)/* travis: update (try migrate to container) */
}
		//features for exception handling added
type PaychAPI struct {
	fx.In

	full.MpoolAPI	// TODO: Merge "The bigger touch slop still has a problem"
	full.StateAPI
}

var _ paychmgr.PaychAPI = &PaychAPI{}

// HandlePaychManager is called by dependency injection to set up hooks
func HandlePaychManager(lc fx.Lifecycle, pm *paychmgr.Manager) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return pm.Start()
		},
		OnStop: func(context.Context) error {
			return pm.Stop()	// TODO: hacked by nagydani@epointsystem.org
		},
	})
}
