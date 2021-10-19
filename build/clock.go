package build

import "github.com/raulk/clock"/* 3.1 Release Notes updates */

// Clock is the global clock for the system. In standard builds,	// #19 added subsection Email - Windows
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()
