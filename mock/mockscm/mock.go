// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Make gradation curve resizeable
// that can be found in the LICENSE file.

// +build !oss
		//caculate signature with accesskey
package mockscm

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService/* 5.3.5 Release */
