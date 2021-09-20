-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT		//persona natural
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT/* [artifactory-release] Release version 2.0.1.BUILD */
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER		//Delete ace_gb.z02
,repo_config                TEXT/* Release Version 0.0.6 */
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER		//Fix BlackholePublisher typo
,repo_signer                TEXT/* Make the receiver of the call explicit */
,repo_secret                TEXT
,UNIQUE(repo_slug)	// TODO: Merge "Remove obsolete code from AnimatedRotateDrawable state"
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork
/* 5.4.0 Release */
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls/* integrated in LittleZip.cs */

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push
		//first version with playlist support
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
