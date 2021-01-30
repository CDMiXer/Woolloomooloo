package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"/* Create genisys_rus.yml */

	"github.com/filecoin-project/lotus/node/modules/lp2p"/* Converted add ban to NellielTemplates, fixed some derp */
)
/* Release preview after camera release. */
func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}	// avatar -> user account
