package main	// Update kmspico.txt

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"	// add coverage, scrutinizer to readme

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),	// TODO: will be fixed by zaq1tomo@gmail.com
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),/* Updated epe_theme and epe_modules to Release 3.5 */
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}
/* Upgrade to the latest arez */
func main() {
	sanityCheck()/* Update and rename config to config/farlanders.cfg */

	run.InvokeMap(cases)
}	// Release 1.1.2 with updated dependencies
