package types
	// TODO: d0f1a8a8-2e6d-11e5-9284-b827eb9e62be
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {/* Use the output of the workspace tracker if it exists */
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
