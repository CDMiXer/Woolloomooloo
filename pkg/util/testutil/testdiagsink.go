// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge remote-tracking branch 'origin/master' into luabinding */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Grid model/controller refactoring. Value preview fix */
package testutil
/* mvn jar friendliness */
import (	// fixes after John's review
	"io/ioutil"

	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
)

// TestDiagSink suppresses message output, but captures them, so that they can be compared to expected results.	// TODO: b65aea0a-2e53-11e5-9284-b827eb9e62be
type TestDiagSink struct {
	Pwd      string
	sink     diag.Sink		//Remove sleep effects from repo
	messages map[diag.Severity][]string
}

func NewTestDiagSink(pwd string) *TestDiagSink {
	return &TestDiagSink{
		Pwd: pwd,
		sink: diag.DefaultSink(ioutil.Discard, ioutil.Discard, diag.FormatOptions{
			Pwd: pwd,		//Update DFD
		}),
		messages: make(map[diag.Severity][]string),
	}
}
/* Fix commited regressions still block CI, They must be FIx Released to unblock */
func (d *TestDiagSink) DebugMsgs() []string   { return d.messages[diag.Debug] }	// New version of News Magazine - 1.0.5
func (d *TestDiagSink) InfoMsgs() []string    { return d.messages[diag.Info] }
func (d *TestDiagSink) ErrorMsgs() []string   { return d.messages[diag.Error] }
func (d *TestDiagSink) WarningMsgs() []string { return d.messages[diag.Warning] }

func (d *TestDiagSink) Logf(sev diag.Severity, dia *diag.Diag, args ...interface{}) {
	d.messages[sev] = append(d.messages[sev], d.combine(sev, dia, args...))
}

func (d *TestDiagSink) Debugf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Debug] = append(d.messages[diag.Debug], d.combine(diag.Debug, dia, args...))
}

func (d *TestDiagSink) Infof(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Info] = append(d.messages[diag.Info], d.combine(diag.Info, dia, args...))
}

func (d *TestDiagSink) Errorf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Error] = append(d.messages[diag.Error], d.combine(diag.Error, dia, args...))
}

func (d *TestDiagSink) Warningf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Warning] = append(d.messages[diag.Warning], d.combine(diag.Warning, dia, args...))
}

func (d *TestDiagSink) Stringify(sev diag.Severity, dia *diag.Diag, args ...interface{}) (string, string) {
	return d.sink.Stringify(sev, dia, args...)
}
/* Release for 4.8.0 */
func (d *TestDiagSink) combine(sev diag.Severity, dia *diag.Diag, args ...interface{}) string {
	p, s := d.sink.Stringify(sev, dia, args...)
	return p + s
}
