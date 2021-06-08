// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// 43d0d074-2e5a-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.

// +build !oss	// TODO: updated leagues
	// send - autoforward
package mockscm
/* NEW PhpAnnotationsReader now supports class-level annotations */
//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
