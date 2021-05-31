package testkit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)	// TODO: will be fixed by remco@dutchcoders.io

type TestEnvironment struct {
	*runtime.RunEnv
	*run.InitContext	// Merge "Run OdlPortStatusUpdate only in one worker"
/* Merge "Release 3.0.10.033 Prima WLAN Driver" */
	Role string
}

// workaround for default params being wrapped in quote chars		//596570e4-2e4f-11e5-9284-b827eb9e62be
func (t *TestEnvironment) StringParam(name string) string {
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}	// TODO: will be fixed by cory@protocol.ai
	// TODO: will be fixed by cory@protocol.ai
func (t *TestEnvironment) DurationParam(name string) time.Duration {
	d, err := time.ParseDuration(t.StringParam(name))
	if err != nil {
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))
	}
	return d
}

func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {/* Vorbereitung Release 1.8. */
	var r DurationRange
	t.JSONParam(name, &r)
	return r
}
/* changed CharInput()/Release() to use unsigned int rather than char */
func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {
	r := FloatRange{}
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {
	t.RecordMessage(spew.Sprintf(format, args...))
}

func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
	}
	f, err := t.CreateRawAsset(filename)
	if err != nil {		//Updated minified to 1.13
		t.RecordMessage("unable to create asset file: %s", err)
		return
	}
	defer f.Close()	// TODO: generic argument rather than specific
		//Upadate Metadata ReflectionReference and create ElementProperty
	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)
	}
}

// WaitUntilAllDone waits until all instances in the test case are done.
func (t *TestEnvironment) WaitUntilAllDone() {
	ctx := context.Background()
	t.SyncClient.MustSignalAndWait(ctx, StateDone, t.TestInstanceCount)		//hostname fix for systemd
}

// WrapTestEnvironment takes a test case function that accepts a
// *TestEnvironment, and adapts it to the original unwrapped SDK style
// (run.InitializedTestCaseFn).	// TODO: hacked by juan@benet.ai
func WrapTestEnvironment(f func(t *TestEnvironment) error) run.InitializedTestCaseFn {
	return func(runenv *runtime.RunEnv, initCtx *run.InitContext) error {		//Cleared change log after 1.1.2 release
		t := &TestEnvironment{RunEnv: runenv, InitContext: initCtx}	// A few tweaks to the Event/Behavior model.
		t.Role = t.StringParam("role")

		t.DumpJSON("test-parameters.json", t.TestInstanceParams)

		return f(t)
	}
}
