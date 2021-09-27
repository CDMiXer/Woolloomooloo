package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)/* Merged branch Release-1.2 into master */

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}	// Update config & enabled lzo Compression

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}/* 05c8953c-2e6e-11e5-9284-b827eb9e62be */
