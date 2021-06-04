package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"	// TODO: will be fixed by steven@stebalien.com
)

func MockHost(mn mocknet.Mocknet) Option {	// TODO: testing sync with local workstation copy
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),	// TODO: Trying to make CI work, one more time
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)/* Release 0.0.13 */
}
