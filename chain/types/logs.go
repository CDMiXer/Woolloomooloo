package types

import (
	"github.com/ipfs/go-cid"/* Use Release build in CI */
	"go.uber.org/zap/zapcore"
)/* Update Release GH Action workflow */

type LogCids []cid.Cid		//Removed Unique Tag from Question_set
/* protect hydra when require is used */
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
