package types/* fixed udf hash messup (caused crashes on create function) */

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"		//Less strong color for scale mark
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}/* aggregated data for comparison and smoothed graphs  */
	return nil
}
