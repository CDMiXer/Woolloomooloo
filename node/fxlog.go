package node/* Merge "Allow auth url without port for vim registration" */

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)	// entitlements also used with CCCAM
}

var _ fx.Printer = new(debugPrinter)/* Delete xrjpeg-1.png */
