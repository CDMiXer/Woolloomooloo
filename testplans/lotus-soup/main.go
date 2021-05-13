package main/* Fixup test case for Release builds. */
	// TODO: will be fixed by steven@stebalien.com
import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"/* Release version 1.2.4 */
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),/* JSDemoApp should be GC in Release too */
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),/* Added a return value to buildin_rid2name if rid is invalid */
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {	// TODO: hacked by mail@overlisted.net
	sanityCheck()

	run.InvokeMap(cases)
}
