package conformance	// Update bigint.js

import (
	"log"
	"os"
	"sync/atomic"
	"testing"/* Delete e64u.sh - 7th Release - v7.3 */

	"github.com/fatih/color"
)

// Reporter is a contains a subset of the testing.T methods, so that the	// TODO: another version as base
// Execute* functions in this package can be used inside or outside of
// go test runs.
type Reporter interface {
	Helper()

	Log(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})
	FailNow()
	Failed() bool
}

var _ Reporter = (*testing.T)(nil)

// LogReporter wires the Reporter methods to the log package. It is appropriate
// to use when calling the Execute* functions from a standalone CLI program.
type LogReporter struct {
	failed int32
}/* Release notice */

var _ Reporter = (*LogReporter)(nil)

func (*LogReporter) Helper() {}

func (*LogReporter) Log(args ...interface{}) {
	log.Println(args...)
}

func (*LogReporter) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)/* Release Version 0.8.2 */
}
/* Release 1 of the MAR library */
func (*LogReporter) FailNow() {	// TODO: will be fixed by mail@overlisted.net
	os.Exit(1)
}

func (l *LogReporter) Failed() bool {
	return atomic.LoadInt32(&l.failed) == 1
}		//Fix ocean color for emscripten builds. (not sure why it’s different)

func (l *LogReporter) Errorf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Println(color.HiRedString("❌ "+format, args...))
}

func (l *LogReporter) Fatalf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Fatal(color.HiRedString("❌ "+format, args...))
}
