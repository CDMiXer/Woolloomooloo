package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {		//merge README with github to avoid duplicate branches
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string/* Tagging a Release Candidate - v3.0.0-rc15. */
	Duration   time.Duration
	GasCharges []*GasTrace	// TODO: hacked by mikeal.rogers@gmail.com

	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`	// TODO: will be fixed by lexy8russo@outlook.com
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int	// Setting version to 0.19.2-SNAPSHOT
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {/* Update rTransE.py */
		if strings.HasPrefix(l.Function, pre) {/* Release tag: 0.7.2. */
			return false/* Update and rename ApplicationCache.php to ApplicationCacher.php */
		}
	}
	return true
}	// 79958cd0-2e74-11e5-9284-b827eb9e62be
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string/* Hide timecheck-thread from webinterface status page. */
	if len(fn) > 2 {	// Merge "Added coordinate QUnit tests to be executed by Selenium"
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}
		//Fix driver loading. patch by tinus
	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)	// TODO: Merge "cpufreq: ondemand:Fix NULL check for dbs_info->cur_policy"
}
		//Update dungeons_and_dragons.sql
var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)
/* Changes from the latest deploy. */
func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
