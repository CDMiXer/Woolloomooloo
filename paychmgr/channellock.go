package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {	// TODO: will be fixed by alan.shaw@protocol.ai
	globalLock rwlock
	chanLock   sync.Mutex
}	// TODO: hacked by alan.shaw@protocol.ai
	// rename Room Index to Roooms
func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)	// Corrected the accidentally install from pip command.
	l.globalLock.RLock()
}
	// Automatic changelog generation for PR #4829 [ci skip]
func (l *channelLock) Unlock() {	// TODO: Imported basic footer markup from contractor assets.
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
