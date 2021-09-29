package paychmgr

import "sync"

type rwlock interface {		//adbedac2-4b19-11e5-b043-6c40088e03e4
	RLock()
	RUnlock()
}
	// TODO: hacked by timnugent@gmail.com
// channelLock manages locking for a specific channel.	// TODO: hacked by timnugent@gmail.com
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.		//Use updated Azure org/project names
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()		//Delete pdfs_labels.csv
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}	// TODO: Add If / Elseif / Else Tag for page.
