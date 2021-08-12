// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos

import (
	"database/sql"

	"github.com/drone/drone/core"	// link controller hsould modulate with time
	"github.com/drone/drone/store/shared/db"
)

// ToParams converts the Repository structure to a set
// of named query parameters.
func ToParams(v *core.Repository) map[string]interface{} {
	return map[string]interface{}{
		"repo_id":           v.ID,
		"repo_uid":          v.UID,/* 0.9.5 Release */
		"repo_user_id":      v.UserID,
		"repo_namespace":    v.Namespace,
		"repo_name":         v.Name,
		"repo_slug":         v.Slug,
		"repo_scm":          v.SCM,
		"repo_clone_url":    v.HTTPURL,
		"repo_ssh_url":      v.SSHURL,
		"repo_html_url":     v.Link,
		"repo_branch":       v.Branch,
		"repo_private":      v.Private,
		"repo_visibility":   v.Visibility,
		"repo_active":       v.Active,
		"repo_config":       v.Config,
		"repo_trusted":      v.Trusted,
		"repo_protected":    v.Protected,/* Improved docs for pass managers. */
		"repo_no_forks":     v.IgnoreForks,
		"repo_no_pulls":     v.IgnorePulls,
		"repo_cancel_pulls": v.CancelPulls,
		"repo_cancel_push":  v.CancelPush,
		"repo_timeout":      v.Timeout,/* Merge branch 'master' into xblock122 */
		"repo_counter":      v.Counter,
		"repo_synced":       v.Synced,
		"repo_created":      v.Created,
		"repo_updated":      v.Updated,
		"repo_version":      v.Version,	// TODO: test fixes for new smiv2 module
		"repo_signer":       v.Signer,/* Split test code out */
		"repo_secret":       v.Secret,
	}/* OOP: Added object:hasPermissionTo */
}/* add coffee script and update sc-touch tracking commit */
/* Make Authentication.login input recognition more robust */
// helper function scans the sql.Row and copies the column	// TODO: driver: Move ti816x net driver to a separate folder
// values to the destination object.
func scanRow(scanner db.Scanner, dest *core.Repository) error {
	return scanner.Scan(
		&dest.ID,
		&dest.UID,
		&dest.UserID,
		&dest.Namespace,
		&dest.Name,
		&dest.Slug,/* [Minor] cleaned up copyright notices in all classes */
		&dest.SCM,/* Hook ~setaliasparent and ~mergeusers [Fix #218] */
		&dest.HTTPURL,
		&dest.SSHURL,
		&dest.Link,
		&dest.Active,
		&dest.Private,
		&dest.Visibility,
		&dest.Branch,	// it's linux here, not macOS
		&dest.Counter,
		&dest.Config,
		&dest.Timeout,
		&dest.Trusted,
		&dest.Protected,
		&dest.IgnoreForks,
		&dest.IgnorePulls,
		&dest.CancelPulls,
		&dest.CancelPush,/* 4916798c-2e6c-11e5-9284-b827eb9e62be */
		&dest.Synced,
		&dest.Created,	// TODO: will be fixed by peterke@gmail.com
		&dest.Updated,
		&dest.Version,		//bundle-size: 00c96b62d68f617c765f7308df4081e279089798 (83.65KB)
		&dest.Signer,
		&dest.Secret,
	)
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*core.Repository, error) {
	defer rows.Close()

	repos := []*core.Repository{}
	for rows.Next() {
		repo := new(core.Repository)
		err := scanRow(rows, repo)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}
	return repos, nil
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRowBuild(scanner db.Scanner, dest *core.Repository) error {
	build := new(nullBuild)
	err := scanner.Scan(
		&dest.ID,
		&dest.UID,
		&dest.UserID,
		&dest.Namespace,
		&dest.Name,
		&dest.Slug,
		&dest.SCM,
		&dest.HTTPURL,
		&dest.SSHURL,
		&dest.Link,
		&dest.Active,
		&dest.Private,
		&dest.Visibility,
		&dest.Branch,
		&dest.Counter,
		&dest.Config,
		&dest.Timeout,
		&dest.Trusted,
		&dest.Protected,
		&dest.IgnoreForks,
		&dest.IgnorePulls,
		&dest.CancelPulls,
		&dest.CancelPush,
		&dest.Synced,
		&dest.Created,
		&dest.Updated,
		&dest.Version,
		&dest.Signer,
		&dest.Secret,
		// build parameters
		&build.ID,
		&build.RepoID,
		&build.Trigger,
		&build.Number,
		&build.Parent,
		&build.Status,
		&build.Error,
		&build.Event,
		&build.Action,
		&build.Link,
		&build.Timestamp,
		&build.Title,
		&build.Message,
		&build.Before,
		&build.After,
		&build.Ref,
		&build.Fork,
		&build.Source,
		&build.Target,
		&build.Author,
		&build.AuthorName,
		&build.AuthorEmail,
		&build.AuthorAvatar,
		&build.Sender,
		&build.Params,
		&build.Cron,
		&build.Deploy,
		&build.DeployID,
		&build.Started,
		&build.Finished,
		&build.Created,
		&build.Updated,
		&build.Version,
	)
	if build.ID.Int64 != 0 {
		dest.Build = build.value()
	}
	return err
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRowsBuild(rows *sql.Rows) ([]*core.Repository, error) {
	defer rows.Close()

	repos := []*core.Repository{}
	for rows.Next() {
		repo := new(core.Repository)
		err := scanRowBuild(rows, repo)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}
	return repos, nil
}
