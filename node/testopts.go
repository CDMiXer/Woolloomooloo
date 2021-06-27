package node	// now its dinamic

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	// TODO: Update emotion headings (#110)
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(/* Release script: be sure to install libcspm before compiling cspmchecker. */
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),/* Merge "wlan: Release 3.2.3.128A" */

		Override(new(lp2p.RawHost), lp2p.MockHost),/* Release notes, NEWS, and quickstart updates for 1.9.2a1. refs #1776 */
		Override(new(mocknet.Mocknet), mn),
	)
}/* Release of eeacms/www-devel:20.4.1 */
