-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT		//working on document structure
,repo_name                  TEXT/* LUTECE-2242 : Refactoring of Search Advanced Parameters management */
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT	// TODO: hacked by mail@bitpshr.net
,repo_ssh_url               TEXT
,repo_html_url              TEXT
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
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;		//put idea panel into menu

-- name: alter-table-repos-add-column-no-pulls		//added notice
	// TODO: Merge branch 'master' into docs-fix-issue-335
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;		//Merge branch 'monday' into Logan-Monday

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;
	// TODO: Update install-alternation
-- name: alter-table-repos-add-column-cancel-push	// TODO: Merge "Specializing x86 range argument copying"

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
