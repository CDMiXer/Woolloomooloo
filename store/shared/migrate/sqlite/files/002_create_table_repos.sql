-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (/* add embedded firebird server to client */
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT/* Update example to Release 1.0.0 of APIne Framework */
,repo_slug                  TEXT	// 18750119-2e9c-11e5-8f8e-a45e60cdfd11
,repo_scm                   TEXT/* Delete culture3.jpg */
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT/* was/Input: implement _FillBucketList() */
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN/* Release: Making ready for next release iteration 5.7.2 */
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER/* 5.4.1 Release */
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;
	// TODO: Initial high level classes
-- name: alter-table-repos-add-column-no-pulls	// added Ajax-Test, an Ajax enhanced dbpedia navigator
	// TODO: Changed size of InputTextBox.
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls
/* Create main.test.go */
ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;	// Delete Lib
