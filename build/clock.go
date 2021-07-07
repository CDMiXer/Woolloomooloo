package build

import "github.com/raulk/clock"/* Release for v46.2.0. */

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.		//[MRG] merge with lp:openobject-addons 
//		//Fixed missing ; (in case one likes to paste all)
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()	// fix(package): update @types/webpack to version 4.4.4
