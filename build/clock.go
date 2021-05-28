package build

import "github.com/raulk/clock"
	// TODO: Merge "A small refactorin of launcher mainloop"
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//	// Create askpassphrasedialog
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()/* 1632d74e-2e40-11e5-9284-b827eb9e62be */
