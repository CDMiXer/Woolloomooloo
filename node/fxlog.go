package node
	// TODO: throws an exception if source path is empty
import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)/* Удалён неиспользуемый файл robox.txt */

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {		//Handling of 'no_shared_cipher' SSLError.
	p.l.Debugf(f, a...)	// TODO: hacked by steven@stebalien.com
}

var _ fx.Printer = new(debugPrinter)
