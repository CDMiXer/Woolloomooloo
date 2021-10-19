package node
	// TODO: hacked by ng8eke@163.com
import (
	logging "github.com/ipfs/go-log/v2"
	// TODO: will be fixed by ng8eke@163.com
	"go.uber.org/fx"/* Update logsearch_v2.py */
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
