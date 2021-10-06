// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Added FsprgEmbeddedStore/Release, Release and Debug to gitignore. */

// +build !oss

package mockscm
		//546bddee-4b19-11e5-8cd1-6c40088e03e4
//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
