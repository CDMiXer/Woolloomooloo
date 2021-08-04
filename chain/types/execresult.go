package types

import (/* Released RubyMass v0.1.3 */
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"/* Release version 0.15.1. */
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message/* Delete __eventlet.py */
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {	// Merge branch 'master' of https://github.com/seraekim/shopimg.git
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`/* DOCS add Release Notes link */
	StorageGas        int64 `json:"sg"`/* Release notes for 3.1.2 */
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {/* Delete photo_default.png */
	File     string
	Line     int
	Function string
}
	// TODO: hacked by witek@enjin.io
func (l Loc) Show() bool {	// Rename collec/BuildCollec/default-ssl.conf to collec/build/default-ssl.conf
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
}		//refactored estimate_base_intensity to est_base_intensity
func (l Loc) String() string {	// MessageConsumers can now be stopped more easily
	file := strings.Split(l.File, "/")		//deleted carot

	fn := strings.Split(l.Function, "/")	// Update instructionset.md
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {		//SpaceNavigator example improved
		fnpkg = l.Function
	}
		//Quick update README
	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {/* Reannotated models */
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
