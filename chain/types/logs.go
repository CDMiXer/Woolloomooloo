package types/* Release 2.1.12 - core data 1.0.2 */
	// // AdminCartRulesController: wording.
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
	// TODO: will be fixed by nicksavers@gmail.com
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {/* readme: fix link after #13 */
		ae.AppendString(c.String())
	}		//Use let !y = x in .. x .. instead of seq in $! and evaluate (#2273)
	return nil
}
