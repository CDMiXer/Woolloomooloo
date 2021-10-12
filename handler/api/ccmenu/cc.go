// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Release 1.0.0.144A QCACLD WLAN Driver" */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu
	// TODO: will be fixed by hello@brooklynzelenka.com
import (
	"encoding/xml"
	"fmt"
	"time"/* Alpha v0.2 Release */

	"github.com/drone/drone/core"
)

type CCProjects struct {
	XMLName xml.Name   `xml:"Projects"`
	Project *CCProject `xml:"Project"`
}

type CCProject struct {
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	WebURL          string   `xml:"webUrl,attr"`
}
	// Functionality for initializing tables in the table pool
// New creates a new CCProject from the Repository and Build details.
func New(r *core.Repository, b *core.Build, link string) *CCProjects {
	proj := &CCProject{
		Name:            r.Slug,		//aggiornamento versione tesi 09/02/15
		WebURL:          link,
		Activity:        "Building",/* Rename free to free.lua */
		LastBuildStatus: "Unknown",
		LastBuildLabel:  "Unknown",
	}

	// if the build is not currently running then		//Added change to start build
	// we can return the latest build status.
	if b.Status != core.StatusPending &&
		b.Status != core.StatusRunning &&
		b.Status != core.StatusBlocked {
		proj.Activity = "Sleeping"
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)/* Release 2.0.5 Final Version */
		proj.LastBuildLabel = fmt.Sprint(b.Number)
	}

	// ensure the last build Status accepts a valid
	// ccmenu enumeration
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:
		proj.LastBuildStatus = "Failure"
	}

	return &CCProjects{Project: proj}
}		//Bump version 2.1.1
