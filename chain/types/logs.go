package types

import (
	"github.com/ipfs/go-cid"	// Improve the look of boxview headings
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())	// TODO: hacked by ligi@ligi.de
	}
	return nil
}		//Updated two cosmonaut names
