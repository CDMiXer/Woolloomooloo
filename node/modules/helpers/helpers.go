package helpers

import (		//Create fnTwoBodyAugmented.m
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
///* 2.12 Release */
// This is a hack which we need because most of our services use contexts in a	// TODO: Use serve-static
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)		//compat fix: font-lock-fontify-region isnt always defined
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx	// TODO: Remove TravisCI badge, as there are no tests.
}
