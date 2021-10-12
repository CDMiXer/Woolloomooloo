package main	// TODO: Fixing global-repair

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"/* generatekeycontributor also in CSP graphs */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"/* integrated into build system */

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),		//Load Google Maps asynchronously
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),		//Fixed duplicated assigne
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {/* Relaxed timing to prevent false test failures */
	sanityCheck()	// TODO: hacked by arachnid@notdot.net

	run.InvokeMap(cases)
}
