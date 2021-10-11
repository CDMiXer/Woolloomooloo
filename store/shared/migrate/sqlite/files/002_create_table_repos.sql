-- name: create-table-repos
		//add attention to README
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT/* Loaded the project */
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT
,repo_slug                  TEXT
,repo_scm                   TEXT/* Release 2.1.11 */
,repo_clone_url             TEXT/* Update collision.py */
,repo_ssh_url               TEXT
,repo_html_url              TEXT/* Update smolyak_evaluate.jl */
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER	// TODO: Update CoveoSimpleRelatedResults.js
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER		//137d53a4-2e5d-11e5-9284-b827eb9e62be
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push
/* Merge "adding v2 support to cinderclient" */
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
