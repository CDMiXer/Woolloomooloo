package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)
/* - Changed for updated exceptions. */
type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
	// TODO: hacked by mail@overlisted.net
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
