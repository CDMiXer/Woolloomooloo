package types/* Added load-save mode and user delay. */

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {		//Added orElse to maybe and added some explanations
	ExitCode exitcode.ExitCode/* Release bzr-1.6rc3 */
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
