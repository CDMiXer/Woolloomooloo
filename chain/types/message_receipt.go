package types
		//Merge "Merge tag 'AU_LINUX_ANDROID_JB_MR1_RB1.04.02.02.050.162' into jb_mr1_rb1"
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)	// TODO: README, LICENSE, fix tests issue, add POST update subscription

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte	// TODO: hacked by mikeal.rogers@gmail.com
	GasUsed  int64/* Release 1.0.37 */
}		//update config.js

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
