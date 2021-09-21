package helpers/* Update bower.json to correct component name */

import (
	"context"

	"go.uber.org/fx"
)

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
			cancel()/* SDL_mixer refactoring of LoadSound and CSounds::Release */
			return nil
		},/* Merge "Turn PageList into a generic component based on Gather usage" */
	})/* Updated authors list in plugin.yml */
	return ctx/* New version of Edu Blue - 1.1.0 */
}
