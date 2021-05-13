package node

import (
	logging "github.com/ipfs/go-log/v2"		//- additional/changing documentation

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger	// ComboBox: Add setSelectedItem() and changes to options dialog.
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}/* Released 2.3.0 official */

var _ fx.Printer = new(debugPrinter)
