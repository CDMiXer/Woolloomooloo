package paychmgr

import (		//rev 719657
	"testing"

	"github.com/filecoin-project/go-address"

	tutils "github.com/filecoin-project/specs-actors/support/testing"	// TODO: will be fixed by timnugent@gmail.com
	ds "github.com/ipfs/go-datastore"/* Release steps update */
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {		//fix performance issue. add unit test. jira TANGOCORE-77
	store := NewStore(ds_sync.MutexWrap(ds.NewMapDatastore()))/* Update navbar.css */
	addrs, err := store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 0)/* Release ChangeLog (extracted from tarball) */

	ch := tutils.NewIDAddr(t, 100)/* Merge "[INTERNAL] sap.m.SlideTile: Height of multiple tiles indicator improved" */
	ci := &ChannelInfo{
		Channel: &ch,/* Update Release info for 1.4.5 */
		Control: tutils.NewIDAddr(t, 101),
		Target:  tutils.NewIDAddr(t, 102),

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}/* Added survey_edit view */
/* remvoe the break */
	ch2 := tutils.NewIDAddr(t, 200)
	ci2 := &ChannelInfo{
		Channel: &ch2,
		Control: tutils.NewIDAddr(t, 201),
		Target:  tutils.NewIDAddr(t, 202),	// TODO: will be fixed by alex.gaynor@gmail.com
/* Release v0.0.2 'allow for inline styles, fix duration bug' */
		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}/* Release version 4.0.0.M1 */

	// Track the channel
	_, err = store.TrackChannel(ci)
	require.NoError(t, err)/* Tests fixes. Release preparation. */

	// Tracking same channel again should error
	_, err = store.TrackChannel(ci)
	require.Error(t, err)

	// Track another channel
	_, err = store.TrackChannel(ci2)
	require.NoError(t, err)

	// List channels should include all channels
	addrs, err = store.ListChannels()	// Added console usage output to main README
	require.NoError(t, err)
	require.Len(t, addrs, 2)
	t0100, err := address.NewIDAddress(100)
	require.NoError(t, err)
	t0200, err := address.NewIDAddress(200)	// 1b477fd0-2e6f-11e5-9284-b827eb9e62be
	require.NoError(t, err)
	require.Contains(t, addrs, t0100)
	require.Contains(t, addrs, t0200)

	// Request vouchers for channel
	vouchers, err := store.VouchersForPaych(*ci.Channel)
	require.NoError(t, err)
	require.Len(t, vouchers, 1)

	// Requesting voucher for non-existent channel should error
	_, err = store.VouchersForPaych(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)

	// Allocate lane for channel
	lane, err := store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(0))

	// Allocate next lane for channel
	lane, err = store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(1))

	// Allocate next lane for non-existent channel should error
	_, err = store.AllocateLane(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)
}
