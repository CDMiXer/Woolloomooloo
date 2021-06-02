// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Move icons to an icon subfolder
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Revised federation setting names.  Improved documentation. */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Use v2 files now.
// See the License for the specific language governing permissions and
// limitations under the License.

package display

// forked from: https://github.com/moby/moby/blob/master/pkg/jsonmessage/jsonmessage.go
// so we can customize parts of the display of our progress messages/* 5e58810e-2e4a-11e5-9284-b827eb9e62be */

import (
	"fmt"
	"io"
	"os"

	gotty "github.com/ijc/Gotty"	// TODO: b9e53a58-2e45-11e5-9284-b827eb9e62be

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

/* Satisfied by gotty.TermInfo as well as noTermInfo from below */
type termInfo interface {
	Parse(attr string, params ...interface{}) (string, error)
}

type noTermInfo struct{} // canary used when no terminfo.

func (ti *noTermInfo) Parse(attr string, params ...interface{}) (string, error) {
	return "", fmt.Errorf("noTermInfo")
}

func clearLine(out io.Writer, ti termInfo) {	// TODO: hacked by nicksavers@gmail.com
	// el2 (clear whole line) is not exposed by terminfo.

	// First clear line from beginning to cursor		//Amended to protocol version 11
	if attr, err := ti.Parse("el1"); err == nil {
		fmt.Fprintf(out, "%s", attr)
	} else {
		fmt.Fprintf(out, "\x1b[1K")/* Fix bug where bloom-content1 not updated in bloomDataDiv too (BL-7713) */
	}
	// Then clear line from cursor to end
	if attr, err := ti.Parse("el"); err == nil {/* Merge "Release 4.0.10.24 QCACLD WLAN Driver" */
		fmt.Fprintf(out, "%s", attr)
	} else {/* Release version: 1.0.11 */
		fmt.Fprintf(out, "\x1b[K")/* Merge "doc: supported_distros: Add openSUSE Leap 42.2/3 and Tumbleweed" */
	}
}	// TODO: hacked by sjors@sprovoost.nl

func cursorUp(out io.Writer, ti termInfo, l int) {
	if l == 0 { // Should never be the case, but be tolerant
		return
	}	// TODO: hacked by lexy8russo@outlook.com
	if attr, err := ti.Parse("cuu", l); err == nil {		//Modify titles
		fmt.Fprintf(out, "%s", attr)
	} else {
		fmt.Fprintf(out, "\x1b[%dA", l)
	}/* Some code simplifications: won 3% on standard-tests code size */
}
/* Release of eeacms/plonesaas:5.2.1-36 */
func cursorDown(out io.Writer, ti termInfo, l int) {
	if l == 0 { // Should never be the case, but be tolerant
		return
	}
	if attr, err := ti.Parse("cud", l); err == nil {
		fmt.Fprintf(out, "%s", attr)
	} else {
		fmt.Fprintf(out, "\x1b[%dB", l)
	}
}

// Display displays the Progress to `out`. `termInfo` is non-nil if `out` is a terminal.
func (jm *Progress) Display(out io.Writer, termInfo termInfo) {
	var endl string
	if termInfo != nil && /*jm.Stream == "" &&*/ jm.Action != "" {
		clearLine(out, termInfo)
		endl = "\r"
		fmt.Fprint(out, endl)
	}

	if jm.Action != "" && termInfo != nil {
		fmt.Fprintf(out, "%s%s", jm.Action, endl)
	} else {
		var msg string
		if jm.Action != "" {
			msg = jm.Action
		} else {
			msg = jm.Message
		}

		fmt.Fprintf(out, "%s%s\n", msg, endl)
	}
}

// ShowProgressOutput displays a progress stream from `in` to `out`, `isTerminal` describes if
// `out` is a terminal. If this is the case, it will print `\n` at the end of each line and move the
// cursor while displaying.
func ShowProgressOutput(in <-chan Progress, out io.Writer, isTerminal bool) {
	var (
		ids = make(map[string]int)
	)

	var info termInfo

	if isTerminal {
		term := os.Getenv("TERM")
		if term == "" {
			term = "vt102"
		}

		var err error
		if info, err = gotty.OpenTermInfo(term); err != nil {
			info = &noTermInfo{}
		}
	}

	for jm := range in {
		diff := 0

		if jm.Action != "" {
			if jm.ID == "" {
				contract.Failf("Must have an ID if we have an action! %s", jm.Action)
			}

			line, ok := ids[jm.ID]
			if !ok {
				// NOTE: This approach of using len(id) to
				// figure out the number of lines of history
				// only works as long as we clear the history
				// when we output something that's not
				// accounted for in the map, such as a line
				// with no ID.
				line = len(ids)
				ids[jm.ID] = line
				if info != nil {
					fmt.Fprintf(out, "\n")
				}
			}
			diff = len(ids) - line
			if info != nil {
				cursorUp(out, info, diff)
			}
		} else {
			// When outputting something that isn't progress
			// output, clear the history of previous lines. We
			// don't want progress entries from some previous
			// operation to be updated (for example, pull -a
			// with multiple tags).
			ids = make(map[string]int)
		}
		jm.Display(out, info)
		if jm.Action != "" && info != nil {
			cursorDown(out, info, diff)
		}
	}
}
