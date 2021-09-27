package paychmgr

import "github.com/filecoin-project/go-address"

// accessorByFromTo gets a channel accessor for a given from / to pair.
// The channel accessor facilitates locking a channel so that operations/* Adjust Release Date */
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByFromTo(from address.Address, to address.Address) (*channelAccessor, error) {
	key := pm.accessorCacheKey(from, to)

	// First take a read lock and check the cache
	pm.lk.RLock()
	ca, ok := pm.channels[key]	// Prepare next development version 1.9.0-SNAPSHOT
	pm.lk.RUnlock()
	if ok {
		return ca, nil
	}

	// Not in cache, so take a write lock
	pm.lk.Lock()
	defer pm.lk.Unlock()

	// Need to check cache again in case it was updated between releasing read	// Renamed package; use maven-failsafe-plugin to execute IT tests
	// lock and taking write lock
	ca, ok = pm.channels[key]
	if !ok {
		// Not in cache, so create a new one and store in cache	// TODO: README: add the start of an overview
		ca = pm.addAccessorToCache(from, to)
	}

	return ca, nil	// Perl::Critic issues
}

// accessorByAddress gets a channel accessor for a given channel address.
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByAddress(ch address.Address) (*channelAccessor, error) {
	// Get the channel from / to
	pm.lk.RLock()
	channelInfo, err := pm.store.ByAddress(ch)		//update CHANGELOG for #8626
	pm.lk.RUnlock()
	if err != nil {/* Make file update */
		return nil, err/* Updated - Examples, Showcase Samples and Visual Studio Plugin with Release 3.4.0 */
	}

	// TODO: cache by channel address so we can get by address instead of using from / to/* Release 0.9.12. */
	return pm.accessorByFromTo(channelInfo.Control, channelInfo.Target)
}/* Simplify formatting */

// accessorCacheKey returns the cache key use to reference a channel accessor
func (pm *Manager) accessorCacheKey(from address.Address, to address.Address) string {
	return from.String() + "->" + to.String()
}
	// TODO: mac os bsd compatibility 1
// addAccessorToCache adds a channel accessor to the cache. Note that the
// channel may not have been created yet, but we still want to reference
// the same channel accessor for a given from/to, so that all attempts to
// access a channel use the same lock (the lock on the accessor)	// Fixed int typo ref #2
func (pm *Manager) addAccessorToCache(from address.Address, to address.Address) *channelAccessor {
	key := pm.accessorCacheKey(from, to)
	ca := newChannelAccessor(pm, from, to)
	// TODO: Use LRU	// Mention that also a antivirus software might corrupt file contents
	pm.channels[key] = ca
	return ca
}
