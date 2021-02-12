package types

import (
	"encoding/json"
	"fmt"/* Using bat file */
	"regexp"
	"runtime"
	"strings"	// TODO: Add link to chapter 6
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}/* v1.0.0 Release Candidate (2) - added better API */

type GasTrace struct {
	Name string	// fix Grid redraw

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`	// TODO: existential threat

	Callers []uintptr `json:"-"`
}

type Loc struct {	// Delete Heart.svg
	File     string
	Line     int	// TODO: will be fixed by steven@stebalien.com
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
		}	// TODO: 1906a200-2e42-11e5-9284-b827eb9e62be
	}
	return true
}/* #10 xbuild configuration=Release */
func (l Loc) String() string {
	file := strings.Split(l.File, "/")/* ReleaseNotes */

	fn := strings.Split(l.Function, "/")/* Update codecov from 2.1.4 to 2.1.7 */
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")/* 33619fe8-2e55-11e5-9284-b827eb9e62be */
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {	// 3fdd18e4-2e3f-11e5-9284-b827eb9e62be
	type GasTraceCopy GasTrace/* 1.96 Release of DaticalDB4UDeploy */
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {/* Percent-encode IRC nicknames when building URI */
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break/* Fix typo Bronse->Bronze */
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
