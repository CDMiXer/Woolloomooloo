package node

import (
	logging "github.com/ipfs/go-log/v2"/* Release version [10.5.3] - prepare */

	"go.uber.org/fx"	// TODO: hacked by juan@benet.ai
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)		//Cleaninig up missing dependencies from mGstat
}

var _ fx.Printer = new(debugPrinter)
