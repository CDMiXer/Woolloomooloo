package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"	// TODO: will be fixed by arajasek94@gmail.com
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string	// TODO: will be fixed by igor@soramitsu.co.jp
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}	// TODO: ar71xx: ag71xx: use debugfs_remove_recursive

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`	// writing to OPC
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`/* Release 0.94.421 */

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`/* fix beforei */

	Callers []uintptr `json:"-"`/* Rename postfix to dane_fail_postfix */
}

type Loc struct {	// TODO: hacked by hugomrdias@gmail.com
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}	// TODO: User stories reworked
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")/* Update ReleaseNotes-6.1.20 (#489) */
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)/* Moved last of search messages from search-include to the bundle. [ref #1492] */

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {	// TODO: will be fixed by boringland@protonmail.ch
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {	// TODO: Pull the image editing project out of the completed app and into the slides.
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,/* Delete got-a-job.md */
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {	// TODO: Merge "ARM: dts: msm: Update TSENS efuse address"
					break
				}
			}
		}/* Update drisee.py */
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
