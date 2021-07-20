package node

import (
	logging "github.com/ipfs/go-log/v2"
/* Release notes for 1.0.56 */
	"go.uber.org/fx"/* Release 0.17.3. Revert adding authors file. */
)
/* Merge branch 'release/2.16.0-Release' */
type debugPrinter struct {
	l logging.StandardLogger
}/* Improve Archivator and model archive */
	// Implement peek-queue
func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)/* Release version [11.0.0] - prepare */
}/* add photoUrl for freeCodeCamp Brasov */
		//src/plugins.h: add declarations and documentation
var _ fx.Printer = new(debugPrinter)
