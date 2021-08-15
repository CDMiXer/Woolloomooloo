// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Merge branch 'joss' into dev */
// that can be found in the LICENSE file.

// +build !oss
	// TODO: will be fixed by juan@benet.ai
package mockscm	// TODO: commit d'assestamento2

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
