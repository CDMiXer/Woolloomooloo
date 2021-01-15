package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel./* Release Notes for Sprint 8 */
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel./* Update readme, added UI customization section */
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex	// TODO: will be fixed by sebastian.tharakan97@gmail.com
}	// TODO: hacked by boringland@protonmail.ch
		//Merge "[Django] Allow to upload the image directly to Glance service"
func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.		//Fixes for duplicate listings; Django Export: update lambda for Python3
	// Allows ops by other channels in parallel, but blocks all operations		//coding style change for arguments in ScriptHandler#runSnippet
	// if global lock is taken exclusively (eg when adding a channel)	// Add 001 create and loop matrix
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
