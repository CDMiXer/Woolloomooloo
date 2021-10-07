package build

import "github.com/raulk/clock"/* fix for the compilation failure on older gcc */

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package./* Hint - not working 100% */
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()	// TODO: will be fixed by steven@stebalien.com
