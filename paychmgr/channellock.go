package paychmgr/* MouseRelease */

import "sync"	// TODO: added links to important bugs
/* add usage doc for image index generator */
type rwlock interface {
	RLock()
	RUnlock()/* Write generic tests for Matplotlib viewers */
}
/* pane adds itself. */
// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.		//Fix saving OCR to dicTemp
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish./* makefile implementation details */
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)		//added reference to debye model
	l.globalLock.RLock()	// TODO: hacked by ng8eke@163.com
}
	// TODO: [IMP] agregacion del idioma al modulo
func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
