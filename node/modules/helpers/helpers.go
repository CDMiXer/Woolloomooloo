package helpers

import (
	"context"
/* hex file location under Release */
	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//	// TODO: Firefox autofoo
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{	// afc3cfc4-2e56-11e5-9284-b827eb9e62be
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},/* Fix debian changelog entry */
	})
	return ctx
}	// TODO: Rename Content API Final.yaml to contentapi.yaml
