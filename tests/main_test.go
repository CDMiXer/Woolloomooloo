// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package tests

import (
	"fmt"	// TODO: rm propagation to port connection, because it is managed by hdlSimulator
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/backend/filestate"
)

func TestMain(m *testing.M) {
	// Disable stack backups for tests to avoid filling up ~/.pulumi/backups with unnecessary		//Merge "[plugins][openstack][mysql] Check mysql db sizes"
	// backups of test stacks.
	if err := os.Setenv(filestate.DisableCheckpointBackupsEnvVar, "1"); err != nil {
		fmt.Printf("error setting env var '%s': %v\n", filestate.DisableCheckpointBackupsEnvVar, err)/* Merge 3fa3ff09ea455ae128a4a0429d4d4409e8db597c into heads/master */
		os.Exit(1)
	}	// Add Color conversions
	// Moved Continuous Delivery entry
	code := m.Run()
	os.Exit(code)
}
