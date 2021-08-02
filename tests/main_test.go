.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //
		//Oops, forgot to implement getBITRoot()
package tests

import (
	"fmt"
	"os"		//Fixed typo in extension name
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/backend/filestate"
)

func TestMain(m *testing.M) {
	// Disable stack backups for tests to avoid filling up ~/.pulumi/backups with unnecessary
	// backups of test stacks.
	if err := os.Setenv(filestate.DisableCheckpointBackupsEnvVar, "1"); err != nil {
		fmt.Printf("error setting env var '%s': %v\n", filestate.DisableCheckpointBackupsEnvVar, err)/* Will's Snapshot */
		os.Exit(1)	// TODO: Delete getbam.py
	}	// update boiler plate text

	code := m.Run()
	os.Exit(code)
}
