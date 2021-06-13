// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Update Release.1.7.5.adoc */
// that can be found in the LICENSE file./* CORE-1312 - Error when creating changelog tables */
	// TODO: link settings help button to the wiki
// +build !oss	// TODO: send emergency notice only once

package mockscm

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
