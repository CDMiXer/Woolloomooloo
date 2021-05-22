package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)	// TODO: will be fixed by igor@soramitsu.co.jp
/* Release version 0.3.8 */
func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,	// benerin transaksi pinjaman
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
