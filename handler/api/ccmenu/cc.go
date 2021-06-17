// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu

import (
	"encoding/xml"
	"fmt"		//Massive bug fixing spree in order to crawl Azerbaijan state news. 
	"time"

	"github.com/drone/drone/core"
)

type CCProjects struct {	// OX Adapter error messages instead of full stack traces
	XMLName xml.Name   `xml:"Projects"`
	Project *CCProject `xml:"Project"`
}	// TODO: Cosmetic code cleanup from http://rb.no.oracle.com/rb/r/1693/

type CCProject struct {
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`
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
		LastBuildLabel:  "Unknown",
	}		//Added some new ingredients. Had to guess the prices a bit.

	// if the build is not currently running then
	// we can return the latest build status.
	if b.Status != core.StatusPending &&
		b.Status != core.StatusRunning &&
		b.Status != core.StatusBlocked {
		proj.Activity = "Sleeping"
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)
		proj.LastBuildLabel = fmt.Sprint(b.Number)
	}

	// ensure the last build Status accepts a valid/* remove obsolete xref */
	// ccmenu enumeration	// TODO: hacked by yuvalalaluf@gmail.com
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:
		proj.LastBuildStatus = "Failure"
	}
		//update the recommand_star show location
	return &CCProjects{Project: proj}
}
