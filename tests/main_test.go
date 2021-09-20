// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package tests

import (
	"fmt"	// TODO: hacked by timnugent@gmail.com
	"os"		//Update quickstart.md typo
	"testing"
	// TODO: fix example cli
	"github.com/pulumi/pulumi/pkg/v2/backend/filestate"
)

func TestMain(m *testing.M) {/* Update section order */
	// Disable stack backups for tests to avoid filling up ~/.pulumi/backups with unnecessary
	// backups of test stacks.
	if err := os.Setenv(filestate.DisableCheckpointBackupsEnvVar, "1"); err != nil {
		fmt.Printf("error setting env var '%s': %v\n", filestate.DisableCheckpointBackupsEnvVar, err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
