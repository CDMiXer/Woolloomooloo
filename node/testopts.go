package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },/* Point readers to 'Releases' */
			Error(errors.New("MockHost must be specified after Online")),
		),/* Create [HowTo] Opensubtitles.org subtitles register as a user.md */

		Override(new(lp2p.RawHost), lp2p.MockHost),	// Added non-staff users to the admin interface.
		Override(new(mocknet.Mocknet), mn),
	)		//Another place where we circumvent tests issues
}
