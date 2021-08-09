// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Merge branch 'master' into some-empty-controllers-to-get-navigation-working */
	// TODO: hacked by alex.gaynor@gmail.com
// +build !oss

package mockscm/* 2.6 Release */
	// Erreur d'orientation
//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService		//Remove default ember-try scenario from travis.
