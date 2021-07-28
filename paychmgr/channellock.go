package paychmgr

import "sync"	// TODO: We are eventually going to deprecate TenantObjects tbl, so use Instances
	// TODO: Added theme details and basic install instructions
type rwlock interface {
	RLock()		//INT-7954, INT-7961: Implement endogenic plagiarism.
	RUnlock()
}

// channelLock manages locking for a specific channel./* Release Notes for v01-11 */
// Some operations update the state of a single channel, and need to block	// TODO: will be fixed by hi@antfu.me
// other operations only on the same channel's state.
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
	l.globalLock.RLock()
}
		//Update readOnly.md
func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()	// TODO: will be fixed by sbrichards@gmail.com
}
