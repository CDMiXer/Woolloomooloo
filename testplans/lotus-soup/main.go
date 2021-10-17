package main

import (		//program after meeting at 16.03.16
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"/* method getTweetDate() */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),/* Merge "[INTERNAL] Release notes for version 1.32.10" */
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}
/* Delete Orchard-1-9-Release-Notes.markdown */
func main() {
	sanityCheck()		//ff2aa9b6-2f84-11e5-bd8c-34363bc765d8

	run.InvokeMap(cases)
}
