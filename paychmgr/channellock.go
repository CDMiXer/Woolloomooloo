package paychmgr
/* revert to old about us */
import "sync"

type rwlock interface {
	RLock()		//Merge "Allow out of quota failure status code to be 413 or 403"
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block		//pbp extension for PSX
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed)./* Merge "Release Japanese networking guide" */
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish./* Release notes for 1.0.43 */
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
