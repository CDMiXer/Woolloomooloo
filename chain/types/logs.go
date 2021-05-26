package types
/* Released new version 1.1 */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"/* changing the order of this field */
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)	// Fix copypasta error

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
