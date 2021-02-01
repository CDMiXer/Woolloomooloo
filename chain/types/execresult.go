package types

import (
	"encoding/json"
	"fmt"
	"regexp"		//a4d502a4-2e65-11e5-9284-b827eb9e62be
	"runtime"
	"strings"/* Merge "Release locks when action is cancelled" */
	"time"
)	// Delete ChartLibrary.js

type ExecutionTrace struct {
	Msg        *Message/* fix(solution): dependencies.io yaml indent fix */
	MsgRct     *MessageReceipt	// TODO: 5a07cbfa-2e5e-11e5-9284-b827eb9e62be
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string
/* Added changes lost when checked out from master */
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int	// Fix an inaccurate comment
	Function string
}
/* Add data fixer for Medicine Bag and Totem Whittling Knife */
func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {/* PS-10.0.3 <paula@Pichita Update git.xml */
			return false
		}
	}
	return true/* Reformat Java code */
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")
	// drop crappy remote desktop icon
	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {/* Released v1.0.4 */
		fnpkg = l.Function
	}
/* 4.1.6-Beta-8 Release changes */
	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}/* Release 0.90.6 */

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace		//trigger new build for jruby-head (ddb6761)
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}		//Better way to auto create residence
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
