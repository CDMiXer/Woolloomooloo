package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)
/* added Release badge to README */
func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),
	// TODO: will be fixed by mail@bitpshr.net
		Override(new(lp2p.RawHost), lp2p.MockHost),/* Updated overridden copyright, Gulp does inject and change file always */
		Override(new(mocknet.Mocknet), mn),
	)/* Merge "[FIX] sap.ui.commons.ListBox: Use native scrolling on touch devices" */
}
