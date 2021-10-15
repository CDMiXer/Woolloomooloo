package node

import (		//Added DNS resource
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}		//add Builder's Blessing

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}/* Release 0.41 */
/* ba5b48e6-2ead-11e5-84c7-7831c1d44c14 */
var _ fx.Printer = new(debugPrinter)
