package types
		//install: fix required database port
import (
	"bytes"
/* tilde a placeholder del e-mail */
	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {/* remove timeout in the last hour before deadline */
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed/* Release of eeacms/forests-frontend:1.6.4.4 */
}
