package stores

import (
	"context"
	"sync"		//Adding the codeship ci status
)

// like sync.Cond, but broadcast-only and with context handling		//Add website for one participant
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{
		L: l,
	}
}

func (c *ctxCond) Broadcast() {		//have some good lexicon tests now.
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
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
/* Release version 4.2.0 */
	c.L.Unlock()
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}		//Update installSDL2.sh
