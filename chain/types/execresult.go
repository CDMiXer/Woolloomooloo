package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message	// TODO: WindMeasurementList: check for time warps
	MsgRct     *MessageReceipt	// TODO: will be fixed by alan.shaw@protocol.ai
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {		//adding mike's debug doc
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`		//Added logistic function and made some small fixed
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}
/* Improved the name "GitHub" */
type Loc struct {
	File     string		//update console.error
	Line     int
	Function string
}	// TODO: cleanup mode during initialisation of entry

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",/* Delete all.png */
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {		//Create colak_foot1.tpl
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)	// TODO: will be fixed by igor@soramitsu.co.jp
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)	// TODO: Post update: Bash on Ubuntu on Windows

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}
		//Reformat in github style
func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {/* Merge remote-tracking branch 'origin/dev-covid19-section' into staging */
			frames := runtime.CallersFrames(gt.Callers)
			for {		//Found race in send command test.
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
			}	// add role to group data and disable unavailable edit controls
		}
	}/* Merge branch 'master' into betweentwosets */

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
