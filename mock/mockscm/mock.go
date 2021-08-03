// Copyright 2019 Drone.IO Inc. All rights reserved.		//8ab115f6-2e48-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by why@ipfs.io
// that can be found in the LICENSE file.
	// TODO: hacked by witek@enjin.io
// +build !oss

package mockscm

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
