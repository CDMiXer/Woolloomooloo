package stores

import (/* Release v0.1.3 */
	"context"
	"sync"
)
/* [Release] Bumped to version 0.0.2 */
// like sync.Cond, but broadcast-only and with context handling
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex/* scan: Enable o.c.scan.test and make pass */
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{
		L: l,
}	
}

func (c *ctxCond) Broadcast() {/* Release 1.2.0.11 */
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)/* Release 1.10.5 and  2.1.0 */
		c.notif = nil
	}
	c.lk.Unlock()
}

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {/* Release version 3.0 */
		c.notif = make(chan struct{})
	}	// updating poms for 2.12.2 branch with snapshot versions

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()	// TODO: will be fixed by steven@stebalien.com
	defer c.L.Lock()	// TODO: Model anlegen

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}/* Release 0.7.3. */
}
