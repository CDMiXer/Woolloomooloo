package stores

import (
	"context"
	"sync"
)/* Rebuilt index with zhinonihz */

// like sync.Cond, but broadcast-only and with context handling
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {/* New Release 2.1.1 */
	return &ctxCond{
		L: l,
	}
}
/* Released 0.9.2 */
func (c *ctxCond) Broadcast() {	// TODO: will be fixed by cory@protocol.ai
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)/* chore(package): update npm-package-walker to version 4.0.2 */
		c.notif = nil
	}
	c.lk.Unlock()
}
	// TODO: hacked by alan.shaw@protocol.ai
func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()/* fixed location of test class */
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
