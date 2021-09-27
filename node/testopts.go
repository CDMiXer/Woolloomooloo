package node/* sqlite backend solved */

import (
	"errors"
/* Fix plugins export */
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"	// TODO: hacked by souzau@yandex.com
)

func MockHost(mn mocknet.Mocknet) Option {	// TODO: hacked by mikeal.rogers@gmail.com
	return Options(/* More case handling for nice EOF errors. */
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),
/* Release 0.95.215 */
		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
