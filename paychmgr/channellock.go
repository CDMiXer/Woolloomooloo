package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel./* Stéphane Maldini */
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block	// TODO: hacked by magik6k@gmail.com
// any operation against any channel.
type channelLock struct {		//Delete parameters.F90
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.	// Created the instance11 for the version1 of the "conference" machine
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()		//SDbShipment
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
