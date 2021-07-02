package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)
/* Update rnaseq-tophat.md */
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Update Release Date for version 2.1.1 at user_guide_src/source/changelog.rst  */
	p.l.Debugf(f, a...)
}
		//Merge pull request #20 from nezhar/master
var _ fx.Printer = new(debugPrinter)
