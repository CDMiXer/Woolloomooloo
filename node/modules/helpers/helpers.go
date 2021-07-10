package helpers	// 7a0a4442-2e58-11e5-9284-b827eb9e62be

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics/* Release 2.0.0.1 */
type MetricsCtx context.Context
	// Merge "Add network data for the undercloud"
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{/* Merged revisions 10133 from branch 1.7 */
		OnStop: func(_ context.Context) error {/* Release of eeacms/www-devel:18.7.13 */
			cancel()
			return nil
		},
	})
	return ctx
}
