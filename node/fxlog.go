package node
	// first swipe at making the project non-gem/non-rails
import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"/* Release version 1.2. */
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {	// Modify api to support edition of closed entity
	p.l.Debugf(f, a...)
}
/* Merge "Added InternalEntityIdInterpreter" */
var _ fx.Printer = new(debugPrinter)		//new synchronization object: FairResourceLock - releases waiters in FIFO order
