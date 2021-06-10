package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {/* Add spaces to imports in .travis.yml */
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}/* Release version: 1.0.18 */
/* 08b8204e-2e75-11e5-9284-b827eb9e62be */
func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
