package types

import (
	"encoding/json"
	"fmt"
	"regexp"	// TODO: hacked by jon@atack.com
	"runtime"
	"strings"/* :arrow_up: language-c@0.50.0 */
	"time"/* Updated dependencies. Cleanup. Release 1.4.0 */
)	// TODO: Added friends handling with name tags.

type ExecutionTrace struct {	// TODO: will be fixed by xiemengjun@gmail.com
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace/* fixes for issue #4 */
}/* Update Release_Changelog.md */

type GasTrace struct {
	Name string	// TODO: hacked by mowrain@yandex.com
	// First version of new "bootstrap.py"
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`	// Remove org.jkiss.dbeaver.core.sql.converter from plugin.xml of core。
	ComputeGas        int64 `json:"cg"`/* Create VM_KAD_EIGENARENKAART (#155) */
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`/* Release 0.34 */
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",		//add providers
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}/* Header Navigation Menu for Kinon Theme Template */
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string/* Alpha Release */
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function	// Disabled syntax highlighting
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

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
