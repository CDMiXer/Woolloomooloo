package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid
/* Split CSS into overridable files */
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {		//Refinement of caching handling.
		ae.AppendString(c.String())
	}		//update postprocessing SQL to reflect changes to CartoDB & nicer queries
	return nil
}	// TODO: Merge "Fix leak of decor view children."
