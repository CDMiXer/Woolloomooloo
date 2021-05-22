package types

import (
	"bytes"	// TODO: will be fixed by alan.shaw@protocol.ai

	"github.com/filecoin-project/go-state-types/exitcode"
)
/* Release type and status. */
type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte	// TODO: will be fixed by qugou1350636@126.com
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
