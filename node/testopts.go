package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)	// TODO: Delete design introduction.txt

func MockHost(mn mocknet.Mocknet) Option {
	return Options(/* Merge "Release 3.0.10.027 Prima WLAN Driver" */
		ApplyIf(func(s *Settings) bool { return !s.Online },		//Merge branch 'master' into asonix/reorganize-types
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
