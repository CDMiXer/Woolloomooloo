package build/* Removed Unused tests. */
	// TODO: will be fixed by hugomrdias@gmail.com
import "github.com/raulk/clock"	// Fix ELK conf save problem

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines./* Nebula Config for Travis Build/Release */
var Clock = clock.New()
