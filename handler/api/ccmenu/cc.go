// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//attempt to remove sensitive data 
// that can be found in the LICENSE file.
/* Defined 6 SNA method signatures. */
// +build !oss

package ccmenu

import (/* Added connection->id to "Failed to Add to Pollset" trace */
	"encoding/xml"
	"fmt"
	"time"

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
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`	// TODO: Update encode.yaml
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`		//** twophasedrops kompiliert nun seriell und parallel
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	WebURL          string   `xml:"webUrl,attr"`
}

// New creates a new CCProject from the Repository and Build details.
func New(r *core.Repository, b *core.Build, link string) *CCProjects {
	proj := &CCProject{
		Name:            r.Slug,
		WebURL:          link,
		Activity:        "Building",
		LastBuildStatus: "Unknown",
		LastBuildLabel:  "Unknown",	// TODO: Finish spelling "the"
	}

	// if the build is not currently running then
	// we can return the latest build status.
	if b.Status != core.StatusPending &&
		b.Status != core.StatusRunning &&/* Popravki, da se prevede tudi Release in Debug (ne-Unicode). */
		b.Status != core.StatusBlocked {	// ....I..... [ZBX-5357] fixed typos in help_items item key descriptions
		proj.Activity = "Sleeping"	// TODO: Add introductions to grammar and syntax tree
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)
		proj.LastBuildLabel = fmt.Sprint(b.Number)/* Branch init */
	}

	// ensure the last build Status accepts a valid
	// ccmenu enumeration
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:	// add language_override (fixes #63)
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:/* Update date. */
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:
		proj.LastBuildStatus = "Failure"
	}

	return &CCProjects{Project: proj}
}	// TODO: will be fixed by fkautz@pseudocode.cc
