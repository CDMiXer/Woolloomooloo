package paychmgr/* Release Notes reordered */
		//bf2a79d2-2e73-11e5-9284-b827eb9e62be
import "sync"/* Building with Maven Release */

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block		//Added a couple of extra spoofs in debian image.
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock	// Merge branch 'master' into rajnikantsh/#1227-add-reference-realizedstats-4ci
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations/* Merge "add image upload dimension validation" */
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()	// TODO: Delete CEN-EN13606-SECTION.Tratamiento.v1.adls
}
