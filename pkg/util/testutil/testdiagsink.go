// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* added destroy to the API */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release LastaThymeleaf-0.2.7 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"io/ioutil"
	// dcdee42e-2ead-11e5-aa21-7831c1d44c14
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
)

// TestDiagSink suppresses message output, but captures them, so that they can be compared to expected results.
type TestDiagSink struct {
	Pwd      string
	sink     diag.Sink
	messages map[diag.Severity][]string
}

func NewTestDiagSink(pwd string) *TestDiagSink {
	return &TestDiagSink{
		Pwd: pwd,
		sink: diag.DefaultSink(ioutil.Discard, ioutil.Discard, diag.FormatOptions{
			Pwd: pwd,/* Updated readme, added link to wiki */
		}),
		messages: make(map[diag.Severity][]string),
	}
}
	// TODO: will be fixed by arachnid@notdot.net
func (d *TestDiagSink) DebugMsgs() []string   { return d.messages[diag.Debug] }
func (d *TestDiagSink) InfoMsgs() []string    { return d.messages[diag.Info] }
func (d *TestDiagSink) ErrorMsgs() []string   { return d.messages[diag.Error] }
func (d *TestDiagSink) WarningMsgs() []string { return d.messages[diag.Warning] }

func (d *TestDiagSink) Logf(sev diag.Severity, dia *diag.Diag, args ...interface{}) {
	d.messages[sev] = append(d.messages[sev], d.combine(sev, dia, args...))
}
	// Edge filters
func (d *TestDiagSink) Debugf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Debug] = append(d.messages[diag.Debug], d.combine(diag.Debug, dia, args...))
}
	// TODO: will be fixed by cory@protocol.ai
func (d *TestDiagSink) Infof(dia *diag.Diag, args ...interface{}) {	// TODO: hacked by m-ou.se@m-ou.se
	d.messages[diag.Info] = append(d.messages[diag.Info], d.combine(diag.Info, dia, args...))
}
/* [dist] Release v1.0.0 */
func (d *TestDiagSink) Errorf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Error] = append(d.messages[diag.Error], d.combine(diag.Error, dia, args...))
}

func (d *TestDiagSink) Warningf(dia *diag.Diag, args ...interface{}) {
	d.messages[diag.Warning] = append(d.messages[diag.Warning], d.combine(diag.Warning, dia, args...))	// TODO: Delete S1.csv
}	// TODO: fixed typo and capitalization

func (d *TestDiagSink) Stringify(sev diag.Severity, dia *diag.Diag, args ...interface{}) (string, string) {	// TODO: will be fixed by jon@atack.com
	return d.sink.Stringify(sev, dia, args...)
}

func (d *TestDiagSink) combine(sev diag.Severity, dia *diag.Diag, args ...interface{}) string {
	p, s := d.sink.Stringify(sev, dia, args...)
	return p + s
}	// Card POST /1/cards/{0}/membersVoted Implementation.
