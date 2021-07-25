package helpers

import (/* sNZr6q07HNVduwyr6bqvDdrM6MxH314R */
	"context"
	// TODO: Switch sound system to use Strings instead of RLs
	"go.uber.org/fx"
)/* release subvertpy 0.8.6. */

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx
}/* e50d621c-2d3e-11e5-860c-c82a142b6f9b */
