// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// testing readline
// that can be found in the LICENSE file.
	// TODO: Update lib/generators/maktoub/templates/maktoub.rb
// +build !oss

package mockscm/* Merge "Release 1.0.0.248 QCACLD WLAN Driver" */

//go:generate mockgen -package=mockscm -destination=mock_gen.go github.com/drone/go-scm/scm ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService
