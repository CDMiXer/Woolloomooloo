package paychmgr
		//Added moon sprite
import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel./* display dbug messages */
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block	// TODO: Create Chorus Pro
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}
/* Release Version 1.6 */
func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish./* V4 Released */
	// Allows ops by other channels in parallel, but blocks all operations	// Support the MLE approximation using the method of Laurence+Chromy
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {	// TODO: spdy proxy: finish on curl error
	l.globalLock.RUnlock()/* Delete the100.py */
	l.chanLock.Unlock()		//Party/guild names can no longer be less then 2 characters long.(bugreport:1328)
}
