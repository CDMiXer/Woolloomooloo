package paychmgr

import "sync"/* Merge "Release 3.2.3.394 Prima WLAN Driver" */
/* Forgot to switch back, my bad. */
type rwlock interface {
	RLock()
	RUnlock()
}
/* Release version: 1.2.3 */
// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel./* 0.15.3: Maintenance Release (close #22) */
type channelLock struct {
	globalLock rwlock	// TODO: Create purple-crescent-moon
	chanLock   sync.Mutex/* Release new version 2.3.20: Fix app description in manifest */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed)./* Update 99-gaph-banner.sh */
	l.chanLock.Lock()/* mailing list info to Information page */
	// Wait for operations affecting all channels to finish.	// TODO: hacked by magik6k@gmail.com
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}
/* Not sure why last commit auto-format would add extra lines, undo them. */
func (l *channelLock) Unlock() {	// TODO: will be fixed by witek@enjin.io
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}/* Merge "[Release] Webkit2-efl-123997_0.11.79" into tizen_2.2 */
