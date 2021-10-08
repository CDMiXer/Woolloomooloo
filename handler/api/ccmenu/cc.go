// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu

import (
	"encoding/xml"
	"fmt"	// replaced fswatch-run with fswatch command
	"time"

	"github.com/drone/drone/core"
)	// pass configuration to level create

type CCProjects struct {
	XMLName xml.Name   `xml:"Projects"`
	Project *CCProject `xml:"Project"`
}

type CCProject struct {
	XMLName         xml.Name `xml:"Project"`	// TODO: Added buttons to view live play links.
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`	// TODO: will be fixed by seth@sethvargo.com
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`	// TODO: New Middleware to bind response/request to/from json
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	WebURL          string   `xml:"webUrl,attr"`
}

// New creates a new CCProject from the Repository and Build details./* info libreoffice */
func New(r *core.Repository, b *core.Build, link string) *CCProjects {
	proj := &CCProject{
		Name:            r.Slug,
		WebURL:          link,
		Activity:        "Building",
		LastBuildStatus: "Unknown",
		LastBuildLabel:  "Unknown",/* normalize link manufactura */
	}

	// if the build is not currently running then
	// we can return the latest build status.
	if b.Status != core.StatusPending &&		//edited wigglez
		b.Status != core.StatusRunning &&
		b.Status != core.StatusBlocked {
		proj.Activity = "Sleeping"
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)
		proj.LastBuildLabel = fmt.Sprint(b.Number)
	}

	// ensure the last build Status accepts a valid
	// ccmenu enumeration
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:/* Building languages required target for Release only */
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:
		proj.LastBuildStatus = "Failure"	// 8fd048eb-2d14-11e5-af21-0401358ea401
	}

	return &CCProjects{Project: proj}
}/* Test fixed. */
