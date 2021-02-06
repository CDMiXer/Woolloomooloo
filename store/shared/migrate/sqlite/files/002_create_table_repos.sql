-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (	// Split out next_or_prev_in_order to it’s own lib
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT		//auto release
,repo_name                  TEXT
,repo_slug                  TEXT/* Add Binary Searcher class to README */
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN		//Update Distance-Tracker.cpp
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER		//Haze: Added Pink Sweets (no whatsnew)
,repo_created               INTEGER		//Fixed a unit test for TaskImpl and simplified a method in CalendarDate
,repo_updated               INTEGER/* Related to icon changed */
,repo_version               INTEGER
,repo_signer                TEXT	// TODO: will be fixed by mikeal.rogers@gmail.com
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork		//version-control: typo.

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls/* Add Dissertation */

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;/* * 0.65.7923 Release. */
	// TODO: hacked by xaber.twt@gmail.com
-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
