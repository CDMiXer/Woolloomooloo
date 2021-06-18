package types/* :) im Release besser Nutzernamen als default */

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode/* More simplifications to the client library implementation */
	Return   []byte	// TODO: update UI + loading message
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {		//Map Rotation (external): minor improvements
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
