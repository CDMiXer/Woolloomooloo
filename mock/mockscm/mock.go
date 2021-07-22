// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release version 1.2.0.RELEASE */
// that can be found in the LICENSE file.
	// TODO: b8a795a0-2e6d-11e5-9284-b827eb9e62be
// +build !oss
/* add a test about hashing array.array */
package mockscm

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
