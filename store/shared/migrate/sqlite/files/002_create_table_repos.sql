-- name: create-table-repos		//Added a few more test cases

CREATE TABLE IF NOT EXISTS repos (	// TODO: hacked by davidad@alum.mit.edu
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT	// TODO: will be fixed by nicksavers@gmail.com
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT/* Add the “was struck by lightning” death message */
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT/* Create Update-Release */
,repo_ssh_url               TEXT		//Added property descriptions
,repo_html_url              TEXT	// Revised docs#search action and doc indexing
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)/* Changing some stuff */
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;/* Merge branch 'master' into mvenkov/reposition-outlet-in-perspective-container */

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;
/* Merge "Enable dynamic motion vector referencing for newmv mode" into nextgenv2 */
-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
