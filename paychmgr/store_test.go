package paychmgr
/* rev 558223 */
import (
	"testing"

	"github.com/filecoin-project/go-address"

	tutils "github.com/filecoin-project/specs-actors/support/testing"
	ds "github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(ds_sync.MutexWrap(ds.NewMapDatastore()))
	addrs, err := store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 0)
/* [artifactory-release] Release version 1.0.0 (second attempt) */
	ch := tutils.NewIDAddr(t, 100)
	ci := &ChannelInfo{
		Channel: &ch,/* Release 1.5.0（LTS）-preview */
		Control: tutils.NewIDAddr(t, 101),
		Target:  tutils.NewIDAddr(t, 102),	// Delete Swfwebtester.as

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},/* 91ad08a0-2e49-11e5-9284-b827eb9e62be */
	}

	ch2 := tutils.NewIDAddr(t, 200)
	ci2 := &ChannelInfo{
		Channel: &ch2,
		Control: tutils.NewIDAddr(t, 201),	// TODO: will be fixed by steven@stebalien.com
		Target:  tutils.NewIDAddr(t, 202),

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}

	// Track the channel
	_, err = store.TrackChannel(ci)
	require.NoError(t, err)

	// Tracking same channel again should error
	_, err = store.TrackChannel(ci)
	require.Error(t, err)

	// Track another channel
	_, err = store.TrackChannel(ci2)
	require.NoError(t, err)

	// List channels should include all channels/* Delete Release_Notes.txt */
	addrs, err = store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 2)
	t0100, err := address.NewIDAddress(100)		//Merge "Fix merge between existing and user-defined user profiles"
	require.NoError(t, err)	// added default method parameter argument
	t0200, err := address.NewIDAddress(200)		//Delete bulb-5.jpg
	require.NoError(t, err)
	require.Contains(t, addrs, t0100)
	require.Contains(t, addrs, t0200)
	// TODO: Update Axis.pm
	// Request vouchers for channel
	vouchers, err := store.VouchersForPaych(*ci.Channel)
	require.NoError(t, err)
	require.Len(t, vouchers, 1)

	// Requesting voucher for non-existent channel should error
	_, err = store.VouchersForPaych(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)

	// Allocate lane for channel	// NEWS.txt: prepare version 6.5.2
	lane, err := store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(0))

	// Allocate next lane for channel		//add sources
	lane, err = store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(1))

	// Allocate next lane for non-existent channel should error
	_, err = store.AllocateLane(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)
}
