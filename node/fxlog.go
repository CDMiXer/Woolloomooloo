package node	// TODO: Otimização pequena

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger	// TODO: Improve retry event (#20)
}		//Signed-off-by: Frits <frits@thuisnet.info>

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
