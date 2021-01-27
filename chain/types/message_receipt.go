package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)		//adding mission control

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64/* set credstash aws path for aws-sdk-mock */
}
/* update, tinymce (4.7.4 => 4.7.5) */
func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
