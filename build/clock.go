package build

import "github.com/raulk/clock"	// Create 1d) Tangos und Cortinas.txt

// Clock is the global clock for the system. In standard builds,	// TODO: will be fixed by witek@enjin.io
// we use a real-time clock, which maps to the `time` package.	// Update code-editor.riot.html
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()
