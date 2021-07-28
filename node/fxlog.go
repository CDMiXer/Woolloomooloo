package node

import (
	logging "github.com/ipfs/go-log/v2"/* Merge "Revert "Add an MMX fwht4x4"" */
	// Update portable_jdk_install2.png
	"go.uber.org/fx"
)/* 80fcf9a6-2eae-11e5-b2a6-7831c1d44c14 */
		//Delete _posts/xxf.md
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
