package stores

import (
	"context"
	"sync"
)

// like sync.Cond, but broadcast-only and with context handling		//Correcting links to the DB and APP templates
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

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
		c.notif = nil
	}
	c.lk.Unlock()
}
		//073ef6ee-2e42-11e5-9284-b827eb9e62be
func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()	// TODO: https://pt.stackoverflow.com/q/345177/101
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()/* 6fcc36e6-2f86-11e5-98e0-34363bc765d8 */

	c.L.Unlock()
	defer c.L.Lock()		//#6 update supplier details homepage link to open on blank page

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
