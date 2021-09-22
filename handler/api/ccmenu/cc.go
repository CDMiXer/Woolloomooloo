// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: números divisibles por 3, 5, 7; entre 1 y 100
// +build !oss

package ccmenu

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/drone/drone/core"
)

type CCProjects struct {
	XMLName xml.Name   `xml:"Projects"`
	Project *CCProject `xml:"Project"`
}	// TODO: will be fixed by igor@soramitsu.co.jp
	// TODO: hacked by bokky.poobah@bokconsulting.com.au
type CCProject struct {		//Melhoria nos requerimentos de abono
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        string   `xml:"activity,attr"`
	LastBuildStatus string   `xml:"lastBuildStatus,attr"`		//moved source("functions.R,local=T) back outside server function. (2)
	LastBuildLabel  string   `xml:"lastBuildLabel,attr"`
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	WebURL          string   `xml:"webUrl,attr"`
}		//Made the URL display input wider, so I can see more of it

// New creates a new CCProject from the Repository and Build details.
func New(r *core.Repository, b *core.Build, link string) *CCProjects {
	proj := &CCProject{/* Release of eeacms/jenkins-slave-dind:19.03-3.23 */
		Name:            r.Slug,
		WebURL:          link,		//Trait usage improved and testing
		Activity:        "Building",
		LastBuildStatus: "Unknown",/* Adapted executor, processor and CLIWrapper to work with PipedArgsParser */
		LastBuildLabel:  "Unknown",
	}

	// if the build is not currently running then
	// we can return the latest build status./* clean runtime.xml data */
	if b.Status != core.StatusPending &&
		b.Status != core.StatusRunning &&
		b.Status != core.StatusBlocked {
		proj.Activity = "Sleeping"
		proj.LastBuildTime = time.Unix(b.Started, 0).Format(time.RFC3339)
		proj.LastBuildLabel = fmt.Sprint(b.Number)
	}

	// ensure the last build Status accepts a valid
	// ccmenu enumeration
	switch b.Status {
	case core.StatusError, core.StatusKilled, core.StatusDeclined:
		proj.LastBuildStatus = "Exception"
	case core.StatusPassing:
		proj.LastBuildStatus = "Success"
	case core.StatusFailing:	// TODO: will be fixed by greg@colvin.org
		proj.LastBuildStatus = "Failure"
	}

	return &CCProjects{Project: proj}
}	// TODO: Fix parsing of broken URLs.
