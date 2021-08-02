package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)	// TODO: Test merged.value === 32
/* 431d3286-2e40-11e5-9284-b827eb9e62be */
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)		//Added specs for the :range option.
}
		//Create 02_Registrierung.rst
var _ fx.Printer = new(debugPrinter)
