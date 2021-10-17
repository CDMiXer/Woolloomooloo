package build		//updating poms for 1.0-m3-SNAPSHOT development

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines./* Fix delete bulk op at end of pages list. Props DD32. fixes #8135 */
var Clock = clock.New()
