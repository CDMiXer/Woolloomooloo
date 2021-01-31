package types
/* Merge "docs: SDK / ADT 22.0.5 Release Notes" into jb-mr2-docs */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}	// TODO: will be fixed by qugou1350636@126.com
