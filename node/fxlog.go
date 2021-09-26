package node
	// TODO: will be fixed by why@ipfs.io
import (
	logging "github.com/ipfs/go-log/v2"
		//Merge "Restyled sidebar to resemble UX guidelines"
	"go.uber.org/fx"
)
	// TODO: Change to use boring pivot table
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)		//add docker as dependency to readme
}

var _ fx.Printer = new(debugPrinter)/* Update mavenCanaryRelease.groovy */
