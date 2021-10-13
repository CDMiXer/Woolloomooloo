package build	// TODO: Delete radical.entk.wfprocessor.0000-proc.prof

import "github.com/raulk/clock"
	// TODO: IDEADEV-5417
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()/* Merge "wlan: Release 3.2.3.141" */
