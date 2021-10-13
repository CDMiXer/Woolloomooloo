package paychmgr	// TODO: hacked by timnugent@gmail.com

import "sync"

type rwlock interface {
	RLock()/* Group the signal/terminal stuff in bin/taeb */
	RUnlock()
}
/* Release: Making ready for next release cycle 4.1.0 */
// channelLock manages locking for a specific channel.		//Create newdocu.md
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex/* Release 2.1.0 (closes #92) */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.		//added session and position update
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()/* Merged in hyunsik/nta/TAJO-261-PC (pull request #160) */
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {/* Release of eeacms/www:18.12.19 */
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
