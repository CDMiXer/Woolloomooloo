package helpers

import (
	"context"

	"go.uber.org/fx"
)	// TODO: will be fixed by sjors@sprovoost.nl

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context
	// Scale down the ant SVG
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {		//Create core
			cancel()
			return nil
		},
	})
	return ctx
}
