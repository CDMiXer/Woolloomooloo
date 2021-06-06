package paychmgr	// TODO: [trunk] Remove old random number functions.
/* More changes to README for singularity 2.4 */
import "sync"/* 3.5 Release Final Release */

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock	// ajout de la ligne pour les nested attributes
	chanLock   sync.Mutex/* Rename footer.html to _includes/footer.html */
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

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()	// TODO: will be fixed by jon@atack.com
	l.chanLock.Unlock()
}		//Added template sfo for xmb icon
