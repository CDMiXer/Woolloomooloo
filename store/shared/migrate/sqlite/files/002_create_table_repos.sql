-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT/* fbdec8d6-2e61-11e5-9284-b827eb9e62be */
,repo_user_id               INTEGER/* -Restore updates GUI which samples were loaded */
,repo_namespace             TEXT
,repo_name                  TEXT
,repo_slug                  TEXT/* Release version 3.0.1 */
,repo_scm                   TEXT	// TODO: hacked by onhardev@bk.ru
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN/* that was silly code */
,repo_private               BOOLEAN		//[at91] tqma9263: update board definition for 3.x series and fix config choose
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER	// TODO: Delete fix.yml
,repo_created               INTEGER/* Rebuilt index with coryreid */
,repo_updated               INTEGER
,repo_version               INTEGER/* Release Notes: update for 4.x */
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork
	// Update SensorMLparsing_IOOSSOS.ipynb
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
