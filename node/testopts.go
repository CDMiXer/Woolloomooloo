package node

import (		//Update link to PRs
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"/* add picture; improve alignment */

	"github.com/filecoin-project/lotus/node/modules/lp2p"	// TODO: hacked by brosner@gmail.com
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),/* Fixing bug with Release and RelWithDebInfo build types. Fixes #32. */
	)
}
