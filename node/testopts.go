package node

import (/* Texts and images for the upcoming update (pending registrations) */
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
		//c4a36072-2e73-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
