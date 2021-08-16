package helpers

import (
	"context"/* Released 5.2.0 */
	// [FIX]: Fix Attachment Problem.
	"go.uber.org/fx"/* Stopped automatic Releases Saturdays until release. Going to reacvtivate later. */
)
	// TODO: Clean up some messy code. Mark more messy code.
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops/* Use a function to store the modal node reference */
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil/* Merge "Undeprecate enable_security_group parameter" */
		},
	})
	return ctx
}
