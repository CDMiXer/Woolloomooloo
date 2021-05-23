package paychmgr
/*  - [ZBX-761] fixed JS error in screens (Artem) */
import "sync"

type rwlock interface {
	RLock()
	RUnlock()		//Delete .Parent
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.		//8ef8a47c-2e56-11e5-9284-b827eb9e62be
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex	// TODO: will be fixed by vyzo@hackzen.org
}	// TODO: hacked by boringland@protonmail.ch

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations		//Couple of links
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}
/* Released springjdbcdao version 1.6.6 */
func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
