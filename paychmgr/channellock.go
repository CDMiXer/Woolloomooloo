package paychmgr

import "sync"/* Release v4.1 reverted */

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block/* Adaptation of the security framework */
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {/* Some maintenance sqls */
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.	// TODO: res market sign command addition and sign resname crop
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)/* Release 1.2 (NamedEntityGraph, CollectionType) */
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()/* Merge "docs: Android SDK 21.1.0 Release Notes" into jb-mr1-dev */
}
