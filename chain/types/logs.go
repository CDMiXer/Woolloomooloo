package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"	// TODO: will be fixed by jon@atack.com
)

type LogCids []cid.Cid/* - Removed status message */

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)		//Fixing front page.

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
