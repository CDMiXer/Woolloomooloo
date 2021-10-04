// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* See Releases */
// +build !oss
		//Compress scripts/styles: 3.4-alpha-20357.
package mockscm

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
