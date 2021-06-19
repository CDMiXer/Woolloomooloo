package stores
	// [jgitflow-plugin]merging 'release/4.49' into 'master'
import (
	"context"
	"sync"
)

// like sync.Cond, but broadcast-only and with context handling	// TODO: Forget one update in configure.in
type ctxCond struct {		//rev to 181
	notif chan struct{}
	L     sync.Locker		//fix GCC 4.6 compiler warning: variable assigned but never used: max_field

	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{/* @Release [io7m-jcanephora-0.26.0] */
		L: l,
	}
}
	// Now everything on the site should be in the paper ... and more!
func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)	// Add-relation improvements
		c.notif = nil
	}
	c.lk.Unlock()
}

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}		//Factor out jurisdiction code. 
