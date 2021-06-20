package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
/* Create imagesfolder */
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {	// * Full app starting, but with layout issues
	return Options(	// c041556a-2e56-11e5-9284-b827eb9e62be
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),/* 0.1.2 Release */
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)		//add new method param by Properties 
}
