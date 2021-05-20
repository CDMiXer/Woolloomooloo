package types	// TODO: will be fixed by seth@sethvargo.com

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
		//Fixed some issues with the nexus to oneliner script
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
