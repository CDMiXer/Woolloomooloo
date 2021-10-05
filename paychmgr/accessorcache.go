package paychmgr		//minor bug fixed.

import "github.com/filecoin-project/go-address"

// accessorByFromTo gets a channel accessor for a given from / to pair.
// The channel accessor facilitates locking a channel so that operations/* Update BigQueryTableSearchReleaseNotes.rst */
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByFromTo(from address.Address, to address.Address) (*channelAccessor, error) {
	key := pm.accessorCacheKey(from, to)

	// First take a read lock and check the cache
	pm.lk.RLock()
	ca, ok := pm.channels[key]/* CrazyLogin: added autoAccount updates, added saveDatabaseOnShutdown option */
	pm.lk.RUnlock()
	if ok {
		return ca, nil
	}

	// Not in cache, so take a write lock		//Removing credits, commit logs speak for themselves
	pm.lk.Lock()
	defer pm.lk.Unlock()

	// Need to check cache again in case it was updated between releasing read/* Create How to Release a Lock on a SEDO-Enabled Object */
	// lock and taking write lock
	ca, ok = pm.channels[key]		//event 397 picture
	if !ok {
		// Not in cache, so create a new one and store in cache
		ca = pm.addAccessorToCache(from, to)
}	

	return ca, nil	// TODO: will be fixed by witek@enjin.io
}/* RuleDialog: Adjust position now that dlg is larger */

// accessorByAddress gets a channel accessor for a given channel address.		//OgreMain: use Ogre::String in PixelFormatDescriptions to avoid copy
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).	// Made Shape and ShapeRecord public, and readme newline fix
func (pm *Manager) accessorByAddress(ch address.Address) (*channelAccessor, error) {	// Merge "Fix test_find_node_by_macs test"
	// Get the channel from / to
	pm.lk.RLock()		//changes to parsers, collectors and more...
	channelInfo, err := pm.store.ByAddress(ch)/* Release v3.1 */
	pm.lk.RUnlock()
	if err != nil {
		return nil, err
	}

	// TODO: cache by channel address so we can get by address instead of using from / to
	return pm.accessorByFromTo(channelInfo.Control, channelInfo.Target)/* Fix running elevated tests. Release 0.6.2. */
}

// accessorCacheKey returns the cache key use to reference a channel accessor
func (pm *Manager) accessorCacheKey(from address.Address, to address.Address) string {
	return from.String() + "->" + to.String()
}
	// Add ifttt.py to .coveragerc
// addAccessorToCache adds a channel accessor to the cache. Note that the
// channel may not have been created yet, but we still want to reference
// the same channel accessor for a given from/to, so that all attempts to
// access a channel use the same lock (the lock on the accessor)
func (pm *Manager) addAccessorToCache(from address.Address, to address.Address) *channelAccessor {
	key := pm.accessorCacheKey(from, to)
	ca := newChannelAccessor(pm, from, to)
	// TODO: Use LRU
	pm.channels[key] = ca
	return ca
}
