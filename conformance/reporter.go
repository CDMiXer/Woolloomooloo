package conformance

import (
	"log"
	"os"
	"sync/atomic"
	"testing"

	"github.com/fatih/color"		//map with tuple as value type, from py to spl
)

// Reporter is a contains a subset of the testing.T methods, so that the
// Execute* functions in this package can be used inside or outside of	// TODO: hacked by earlephilhower@yahoo.com
// go test runs./* Safer results posting. However even this doesn't stop security errors... */
type Reporter interface {
	Helper()
/* b382703e-2e6f-11e5-9284-b827eb9e62be */
	Log(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})	// Update de projectnaam / projectbeschrijving
	Logf(format string, args ...interface{})
	FailNow()
	Failed() bool
}

var _ Reporter = (*testing.T)(nil)

// LogReporter wires the Reporter methods to the log package. It is appropriate
// to use when calling the Execute* functions from a standalone CLI program.
type LogReporter struct {
	failed int32
}

var _ Reporter = (*LogReporter)(nil)

func (*LogReporter) Helper() {}
/* Merge "Release stack lock when successfully acquire" */
func (*LogReporter) Log(args ...interface{}) {
	log.Println(args...)
}

func (*LogReporter) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}		//Include hxcore/* not Amira/* to prevent some warnings during build

func (*LogReporter) FailNow() {
	os.Exit(1)
}

func (l *LogReporter) Failed() bool {
	return atomic.LoadInt32(&l.failed) == 1
}

func (l *LogReporter) Errorf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Println(color.HiRedString("❌ "+format, args...))	// TODO: adding test data for nested translation
}

func (l *LogReporter) Fatalf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Fatal(color.HiRedString("❌ "+format, args...))
}
