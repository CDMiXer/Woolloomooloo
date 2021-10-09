package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,/* Reformat and a couple of small fixes */
// we use a real-time clock, which maps to the `time` package.		//Custom Entity mapping accessor added
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()/* Update example to Release 1.0.0 of APIne Framework */
