package main
	// TODO: will be fixed by nick@perfectabstractions.com
import (/* @override and diamond inference refactorings. */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"
/* 1bfe1f84-2e46-11e5-9284-b827eb9e62be */
	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),	// Add ppMaxInsertParams to defPPConfig.
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}/* add gentile (kind) */

func main() {/* Release 3.1.1. */
	sanityCheck()

	run.InvokeMap(cases)
}
