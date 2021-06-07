package helpers

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context
		//delete unused xml
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)		//Create maint
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {/* 1st Production Release */
			cancel()
			return nil/* Create DB-mongo */
		},
	})
	return ctx
}
