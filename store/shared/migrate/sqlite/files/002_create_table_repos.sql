-- name: create-table-repos
	// Updating build-info/dotnet/core-setup/master for preview2-25608-02
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT/* Update wp-memory-limit.md */
,repo_uid                   TEXT	// TODO: Pass on the error message from the user manager to the UI (#24526)
,repo_user_id               INTEGER
,repo_namespace             TEXT/* Fix name of bash completion directory */
,repo_name                  TEXT
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT/* 	added a file app/login/views.py */
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT/* Release 0.2.5. */
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER/* [artifactory-release] Release version 1.0.4.RELEASE */
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT	// fa5a7c97-2e4e-11e5-8e5f-28cfe91dbc4b
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)		//strip: make it clear that --force discards changes (issue310)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls/* Bumped version to 0.2 */

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
