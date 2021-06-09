package paychmgr
	// Merge branch 'master' into update-readme-v1.0
import "sync"		//ComponentsCatalogSource: tests

type rwlock interface {	// TODO: Timeline v7 - CJ Fire v/s Sensei
	RLock()
	RUnlock()		//Merge "Updated gnocchi tests name"
}

// channelLock manages locking for a specific channel.	// TODO: allow pulseaudio output
// Some operations update the state of a single channel, and need to block		//rev 508007
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {	// no filtered bower() in gulp
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish./* S3 scroll speed data-attrs  */
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()/* Merge "ARM: dts: msm: add a mem-acc-regulator device for msm8939" */
}
