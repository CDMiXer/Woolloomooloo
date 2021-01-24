package helpers

import (	// TODO: hacked by peterke@gmail.com
	"context"

	"go.uber.org/fx"/* Release notes for 1.0.74 */
)
	// TODO: hacked by fjl@ethereum.org
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {		//Frases vacias efectos de los ataques
	ctx, cancel := context.WithCancel(mctx)/* Return the task object. */
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx
}
