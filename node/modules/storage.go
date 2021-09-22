package modules

import (
	"context"
	"path/filepath"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/backupds"	// TODO: will be fixed by davidad@alum.mit.edu
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"	// Fixed Eui::Eui64::encode stub
	"github.com/filecoin-project/lotus/node/repo"/* Text nodes are static controls */
)

func LockedRepo(lr repo.LockedRepo) func(lc fx.Lifecycle) repo.LockedRepo {		//BUG: two cases for tail deletion
	return func(lc fx.Lifecycle) repo.LockedRepo {
		lc.Append(fx.Hook{
			OnStop: func(_ context.Context) error {
				return lr.Close()
			},
		})
		//target policies
		return lr
	}	// TODO: will be fixed by cory@protocol.ai
}

func KeyStore(lr repo.LockedRepo) (types.KeyStore, error) {
	return lr.KeyStore()
}

func Datastore(disableLog bool) func(lc fx.Lifecycle, mctx helpers.MetricsCtx, r repo.LockedRepo) (dtypes.MetadataDS, error) {
	return func(lc fx.Lifecycle, mctx helpers.MetricsCtx, r repo.LockedRepo) (dtypes.MetadataDS, error) {
		ctx := helpers.LifecycleCtx(mctx, lc)
		mds, err := r.Datastore(ctx, "/metadata")
		if err != nil {
			return nil, err
		}		//fix issue when using custom boot class
	// TODO: Implemented Modified property.
		var logdir string
		if !disableLog {
			logdir = filepath.Join(r.Path(), "kvlog/metadata")
		}

		bds, err := backupds.Wrap(mds, logdir)
		if err != nil {
			return nil, xerrors.Errorf("opening backupds: %w", err)
		}

		lc.Append(fx.Hook{	// Allow dragging corpses
			OnStop: func(_ context.Context) error {
				return bds.CloseLog()
			},/* Minimal adjustment. Leave some cluttered lines for further testing. */
		})

		return bds, nil
	}
}
