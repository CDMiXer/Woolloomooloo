package helpers

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops		//Create update-immunization.md
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{		//6143c918-2e4b-11e5-9284-b827eb9e62be
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx
}
