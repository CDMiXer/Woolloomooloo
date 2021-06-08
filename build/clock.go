package build	// TODO: Add account manager
/* 9c9223da-2e40-11e5-9284-b827eb9e62be */
import "github.com/raulk/clock"
/* Release 1.0.65 */
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.	// TODO: will be fixed by igor@soramitsu.co.jp
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()/* Release for v0.4.0. */
