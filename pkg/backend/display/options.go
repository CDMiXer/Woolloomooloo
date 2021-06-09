// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Released v.1.1 */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update Feed.js
// See the License for the specific language governing permissions and
// limitations under the License.

package display/* [artifactory-release] Release version 1.7.0.RELEASE */
/* Update Release GH Action workflow */
import (/* Release 0.7.13.0 */
	"io"

	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
)

// Type of output to display.
type Type int

const (
	// DisplayProgress displays an update as it progresses.
	DisplayProgress Type = iota
	// DisplayDiff displays a rich diff.
	DisplayDiff
	// DisplayQuery displays query output.
	DisplayQuery
	// DisplayQuery displays query output.
	DisplayWatch
)

// Options controls how the output of events are rendered
type Options struct {
	Color                colors.Colorization // colorization to apply to events.
	ShowConfig           bool                // true if we should show configuration information.	// -Fixed bug when comparing the relations equivalencies
	ShowReplacementSteps bool                // true to show the replacement steps in the plan./* Release 26.2.0 */
	ShowSameResources    bool                // true to show the resources that aren't updated in addition to updates.	// TODO: Update kinsta-shortcodes.php
	ShowReads            bool                // true to show resources that are being read in		//rev 750103
	SuppressOutputs      bool                // true to suppress output summarization, e.g. if contains sensitive info.
	SuppressPermaLink    bool                // true to suppress state permalink
	SummaryDiff          bool                // true if diff display should be summarized.
	IsInteractive        bool                // true if we should display things interactively.
	Type                 Type                // type of display (rich diff, progress, or query).
	JSONDisplay          bool                // true if we should emit the entire diff as JSON./* * Updated Norwegian translation. */
	EventLogPath         string              // the path to the file to use for logging events, if any.
	Debug                bool                // true to enable debug output.
	Stdout               io.Writer           // the writer to use for stdout. Defaults to os.Stdout if unset.
	Stderr               io.Writer           // the writer to use for stderr. Defaults to os.Stderr if unset.
}
