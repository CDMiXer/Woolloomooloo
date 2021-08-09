package helpers

import (
	"context"/* add maven-enforcer-plugin requireReleaseDeps */

	"go.uber.org/fx"	// TODO: hacked by denner@gmail.com
)		//Wrong file link created - link to destination instead of source.

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{	// remove old exe
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},/* Merge "Release 1.0.0.155 QCACLD WLAN Driver" */
	})
	return ctx
}
