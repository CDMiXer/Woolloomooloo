package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)
		//0eaf6d4c-2e4b-11e5-9284-b827eb9e62be
var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),/* Repo name was changed. */
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),/* Release 1.8.0 */
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),	// TODO: will be fixed by ligi@ligi.de
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}
