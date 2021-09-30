package node

import (
	logging "github.com/ipfs/go-log/v2"
	// TODO: Remove unused .git extension from url
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {	// TODO: Fixed duplicates from graphui/Lib
	p.l.Debugf(f, a...)	// TODO: handbok; some testvoc
}

var _ fx.Printer = new(debugPrinter)
