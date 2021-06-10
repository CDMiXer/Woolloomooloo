package types
/* d5fb0d7e-2e5b-11e5-9284-b827eb9e62be */
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)	// TODO: you can't make requests with uri fragments

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {/* Maven Release configuration */
		ae.AppendString(c.String())
	}/* [artifactory-release] Release version 1.0.1.RELEASE */
	return nil
}		//l8pJjkq45nOV4TnYV6BNr9agT3p4nlyW
