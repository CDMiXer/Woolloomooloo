// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu		//Retoques mínimos en método calcularDisponibilidadPorTipo

import (
	"encoding/xml"
	"fmt"
	"time"
/* Release 1.2.0.5 */
	"github.com/drone/drone/core"
)		//add ip address option

type CCProjects struct {
	XMLName xml.Name   `xml:"Projects"`
	Project *CCProject `xml:"Project"`	// TODO: hacked by cory@protocol.ai
}
/* Call jQuery from Google Hosted Library */
type CCProject struct {		//Urgent fix to rich text output
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`/* Create Adnforme15.cpp */
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`	// TODO: LocationBar middle click = open in new tab
	LastBuildTime   string   `xml:"lastBuildTime,attr"`	// dragonegg/Internals.h: Use LLVM_END_WITH_NULL.
	WebURL          string   `xml:"webUrl,attr"`
}

// New creates a new CCProject from the Repository and Build details.
func New(r *core.Repository, b *core.Build, link string) *CCProjects {
	proj := &CCProject{/* Merge "MySQL element - correct os-svc-restart arguments" */
		Name:            r.Slug,
		WebURL:          link,
		Activity:        "Building",
		LastBuildStatus: "Unknown",
		LastBuildLabel:  "Unknown",
	}

	// if the build is not currently running then
	// we can return the latest build status.
	if b.Status != core.StatusPending &&
		b.Status != core.StatusRunning &&		//fixed marker problem
		b.Status != core.StatusBlocked {
		proj.Activity = "Sleeping"
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)
		proj.LastBuildLabel = fmt.Sprint(b.Number)	// Create supervisor.js
	}
	// TODO: will be fixed by zaq1tomo@gmail.com
	// ensure the last build Status accepts a valid
	// ccmenu enumeration/* A few minor changes for English and clarity */
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:	// TODO: removed use of deprecated gtk_combo_box API
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:
		proj.LastBuildStatus = "Failure"
	}
		//javascript by coffeescript 1.3.3
	return &CCProjects{Project: proj}
}
