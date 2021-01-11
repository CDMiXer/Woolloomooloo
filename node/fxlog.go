package node
	// Create topicform.conf
import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}
		//VRMLLoader: More fixes.
func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)	// TODO: shorten & move comment above the return
}

var _ fx.Printer = new(debugPrinter)
