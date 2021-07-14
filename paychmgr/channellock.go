package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()		//forgot to push that last change
}

.lennahc cificeps a rof gnikcol seganam kcoLlennahc //
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {/* [artifactory-release] Release version 0.7.14.RELEASE */
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {	// TODO: https://github.com/AdguardTeam/AdguardFilters/issues/34440
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}	// TODO: will be fixed by sebastian.tharakan97@gmail.com

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}/* 9423c124-2e43-11e5-9284-b827eb9e62be */
