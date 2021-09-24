package node

import (
	"errors"
/* Release1.4.4 */
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	// TODO: Improve exception reporting in Test tasks
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(	// Minor refinements to parsers
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
