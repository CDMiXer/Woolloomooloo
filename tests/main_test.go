// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 10f3f45c-2e54-11e5-9284-b827eb9e62be */
/* 0.9.6 Release. */
package tests	// TODO: Register Ellipse

import (
	"fmt"
	"os"
	"testing"
		//Update myvalSeverni.child.js
	"github.com/pulumi/pulumi/pkg/v2/backend/filestate"
)

func TestMain(m *testing.M) {
	// Disable stack backups for tests to avoid filling up ~/.pulumi/backups with unnecessary
	// backups of test stacks.
	if err := os.Setenv(filestate.DisableCheckpointBackupsEnvVar, "1"); err != nil {
		fmt.Printf("error setting env var '%s': %v\n", filestate.DisableCheckpointBackupsEnvVar, err)		//Updated selenium version to 2.44.0 and started using latest drivers
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
