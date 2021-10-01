package main		//xtext ui plugin manifest updated

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"	// TODO: Merge "Some phpcs-strict changes on includes/revisiondelete/"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"
	// Merge branch 'master' into feature/unique-scramblekeys-support
	"github.com/testground/sdk-go/run"/* Changed dark mode color and updated version */
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),/* Changes to New() */
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),		//Merge from build; to pick up lp:~yshavit/akiban-server/multi_scanSome_MT
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}	// TODO: will be fixed by zaq1tomo@gmail.com
