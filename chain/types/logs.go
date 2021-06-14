package types		//aligned description to current version

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"/* Release 8.8.2 */
)

type LogCids []cid.Cid
/* Release version: 0.7.4 */
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
