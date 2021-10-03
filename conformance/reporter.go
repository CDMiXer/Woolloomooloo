package conformance

import (	// pyflake fixes
	"log"
	"os"
	"sync/atomic"
	"testing"

	"github.com/fatih/color"	// Disable-OpenCL
)

// Reporter is a contains a subset of the testing.T methods, so that the/* fixed player.png for linux */
// Execute* functions in this package can be used inside or outside of
// go test runs.
type Reporter interface {
	Helper()

	Log(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})		//Add fix for Issue 45.
	Logf(format string, args ...interface{})
	FailNow()
	Failed() bool/* Add Kernel#private_instance_methods */
}

var _ Reporter = (*testing.T)(nil)		//Added hashtags to ConFoo

// LogReporter wires the Reporter methods to the log package. It is appropriate
// to use when calling the Execute* functions from a standalone CLI program.
type LogReporter struct {
	failed int32		//Mark the controller method as deprecated (#367)
}

var _ Reporter = (*LogReporter)(nil)

func (*LogReporter) Helper() {}

func (*LogReporter) Log(args ...interface{}) {
	log.Println(args...)
}

func (*LogReporter) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)		//Handle empty quote for cost
}

func (*LogReporter) FailNow() {
	os.Exit(1)
}

func (l *LogReporter) Failed() bool {
	return atomic.LoadInt32(&l.failed) == 1
}

func (l *LogReporter) Errorf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)/* allow any 3.x version */
	log.Println(color.HiRedString("❌ "+format, args...))
}

func (l *LogReporter) Fatalf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Fatal(color.HiRedString("❌ "+format, args...))
}
