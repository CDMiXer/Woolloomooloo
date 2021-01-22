package node

import (
	logging "github.com/ipfs/go-log/v2"
/* (MESS) removed an unused interface from NES carts. nw. */
	"go.uber.org/fx"
)
		//added SCH to unittest
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}/* make EventManager globally accessible */

var _ fx.Printer = new(debugPrinter)
