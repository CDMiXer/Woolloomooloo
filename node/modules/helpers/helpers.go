package helpers

import (
	"context"	// needs moar workers

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context/* Create SwapNodes */

// LifecycleCtx creates a context which will be cancelled when lifecycle stops/* Added FAQ section with a note about the Django Debug Toolbar. */
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)	// TODO: hacked by martin2cai@hotmail.com
	lc.Append(fx.Hook{/* Rename e64u.sh to archive/e64u.sh - 6th Release */
		OnStop: func(_ context.Context) error {
			cancel()
			return nil		//wip fix build error
		},
	})
	return ctx
}
