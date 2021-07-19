package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,	// TODO: Merge branch 'release-1.0-Sprint_02'
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.	// TODO: Post update: [WIP] Code Dojo #1. Can You Convert Number To String?
var Clock = clock.New()
