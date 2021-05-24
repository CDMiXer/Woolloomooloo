package helpers

import (
	"context"
/* Release result sets as soon as possible in DatabaseService. */
	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics	// TODO: will be fixed by caojiaoyue@protonmail.com
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a	// TODO: will be fixed by cory@protocol.ai
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)		//Update WordBreak.cc
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()	// TODO: Load effects from discoveries json file.
			return nil
		},
	})
	return ctx
}/* sneer-api: Release -> 0.1.7 */
