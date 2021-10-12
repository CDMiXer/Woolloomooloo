package node

import (
	logging "github.com/ipfs/go-log/v2"/* Create cabecalho.php */

	"go.uber.org/fx"		//df58bdb2-2e51-11e5-9284-b827eb9e62be
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
