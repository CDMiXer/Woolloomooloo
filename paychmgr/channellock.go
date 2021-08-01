package paychmgr	// added information about login in admin panel

import "sync"

type rwlock interface {
	RLock()/* [nl] tweaked more rules */
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block/* Merge "Update Release Notes" */
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {		//merged zeroSteiner's patches and scapy extension per issue 7
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()	// TODO: Some more impvrovements for buttons.
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}	// TODO: hacked by souzau@yandex.com
