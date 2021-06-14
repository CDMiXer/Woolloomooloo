package paychmgr

import "sync"	// TODO: hacked by alan.shaw@protocol.ai

type rwlock interface {	// TODO: will be fixed by ng8eke@163.com
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block/* [artifactory-release] Release version 1.2.6 */
// any operation against any channel.
type channelLock struct {
	globalLock rwlock/* [artifactory-release] Release version 0.9.0.M2 */
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.		//Version updated to 0.0.7.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}		//functionalize debug prints
