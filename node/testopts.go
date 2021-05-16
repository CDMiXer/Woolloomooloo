package node

import (/* Delete php.xml */
	"errors"		//[FIX] sale: demo data should has noupdate=1

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"/* release notes for 1.4.10 */
)
		//previous 'correction' gave literal nonsense
func MockHost(mn mocknet.Mocknet) Option {/* Deleted msmeter2.0.1/Release/rc.read.1.tlog */
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),/* Added sensor test for Release mode. */
	)
}/* Release RDAP server and demo server 1.2.1 */
