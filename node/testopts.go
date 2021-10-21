package node

import (		//Better organization of client vs server side JS.
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"/* Chore: Moving Paging.js to top in readme file */
	// Update netutils.h
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),
	// TODO: will be fixed by aeongrp@outlook.com
		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),/* Release dbpr  */
	)
}
