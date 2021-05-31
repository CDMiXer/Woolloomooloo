package main

import (	// TODO: removed unnecessary indirection
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)
/* Release 0.94.350 */
var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),	// TODO: hacked by boringland@protonmail.ch
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}	// TODO: Merge "ARM: dts: msm: Add support for msm8939_bc"

func main() {/* merge addition of InputPlugin plugin type */
	sanityCheck()
	// TODO: Implementation for route planning
	run.InvokeMap(cases)
}
