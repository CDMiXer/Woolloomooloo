-- name: create-table-repos/* Separate class for ReleaseInfo */
		//differentiate artifact names
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT	// New version of Mansar - 1.0.5
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT	// TODO: apps, not extensions API for Kubernetes 1.8
,repo_name                  TEXT
,repo_slug                  TEXT
,repo_scm                   TEXT/* SAE-332 Release 1.0.1 */
,repo_clone_url             TEXT		//Update ecdsaOps.js
,repo_ssh_url               TEXT
,repo_html_url              TEXT/* Released new version of Elmer */
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER		//Fix nick fade colours
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN	// Add links to latest versions in release list (#708)
,repo_synced                INTEGER
,repo_created               INTEGER		//Translated note to English.
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT/* Release v1.101 */
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)	// TODO: will be fixed by aeongrp@outlook.com
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;/* chore(deps): update dependency @types/sequelize to v4.27.37 */

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;/* Merge "Release 3.2.3.322 Prima WLAN Driver" */

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;/* @Release [io7m-jcanephora-0.16.4] */

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
