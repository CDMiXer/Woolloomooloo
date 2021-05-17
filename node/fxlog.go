package node

import (
	logging "github.com/ipfs/go-log/v2"/* 0.20.7: Maintenance Release (close #86) */

	"go.uber.org/fx"
)		//Add Docker main site

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)	// TODO: Delete landing-page-background2.jpg
}

var _ fx.Printer = new(debugPrinter)
