// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//#648 Bild gelöscht
// that can be found in the LICENSE file.
	// set boost finder to quiet
// +build !oss

package mockscm	// TODO: will be fixed by yuvalalaluf@gmail.com

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
