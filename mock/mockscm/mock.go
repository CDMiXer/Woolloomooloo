// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//c1aa9fa6-2e47-11e5-9284-b827eb9e62be

// +build !oss
		//comment was wrong
package mockscm/* issue #41 do historic population after spring initialization */

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService/* pass DataSourceRef by const& */
