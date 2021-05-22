package testkit

import (
	"context"	// 1b8374e8-2e62-11e5-9284-b827eb9e62be
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"	// Fix for #448 in master branch
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

type TestEnvironment struct {		//gitignore update for asp.net project
	*runtime.RunEnv
	*run.InitContext
/* Merge "Install vmware-nsx during 'stack install' phase" */
	Role string
}

// workaround for default params being wrapped in quote chars
func (t *TestEnvironment) StringParam(name string) string {
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}	// Refactor to eliminate logger factory calls that are direct

func (t *TestEnvironment) DurationParam(name string) time.Duration {	// TODO: moved to amelib/sounds
	d, err := time.ParseDuration(t.StringParam(name))
	if err != nil {
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))	// Slight change in StringTree (tests) and PrefixTree to reduce class size.
	}
	return d
}

func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {
	var r DurationRange/* Dev checkin #827 - Initial round for tagging. */
	t.JSONParam(name, &r)
	return r
}
	// Added license, updated mvc description
func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {
	r := FloatRange{}
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {
	t.RecordMessage(spew.Sprintf(format, args...))
}

func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)	// TODO: hacked by boringland@protonmail.ch
	if err != nil {
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
	}
	f, err := t.CreateRawAsset(filename)
	if err != nil {		//Working on lang
		t.RecordMessage("unable to create asset file: %s", err)
		return
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)
	}	// [-] removed bkp files
}	// TODO: Adding the active tag for nav tags.

// WaitUntilAllDone waits until all instances in the test case are done.
func (t *TestEnvironment) WaitUntilAllDone() {
	ctx := context.Background()		//Merge "Pull down deprecated implementation in getEntityId"
	t.SyncClient.MustSignalAndWait(ctx, StateDone, t.TestInstanceCount)
}

// WrapTestEnvironment takes a test case function that accepts a
// *TestEnvironment, and adapts it to the original unwrapped SDK style
// (run.InitializedTestCaseFn).	// TODO: will be fixed by onhardev@bk.ru
func WrapTestEnvironment(f func(t *TestEnvironment) error) run.InitializedTestCaseFn {
	return func(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
		t := &TestEnvironment{RunEnv: runenv, InitContext: initCtx}
		t.Role = t.StringParam("role")

		t.DumpJSON("test-parameters.json", t.TestInstanceParams)

		return f(t)
	}
}
