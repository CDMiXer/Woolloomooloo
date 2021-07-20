-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT		//Design updating
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT/* Release for 2.14.0 */
,repo_slug                  TEXT
,repo_scm                   TEXT/* Creating branch for update to latest Spring and Hibername */
,repo_clone_url             TEXT/* Adding ReleaseProcess doc */
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN		//FileVersions - latest() implemented
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT/* Modify lifecycle settings */
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN		//Rename ~chairman.html to index.html
,repo_synced                INTEGER
,repo_created               INTEGER	// TODO: [ELF] Add missing -target option.
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);/* bugfix tarakey auf 2 setzen, wenn tara läuft */
		//Merge "Support for use outside of DrawerLayout" into mnc-ub-dev
-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls		//radio buttons for everyone! now with indent...yay...
	// Préparation marshaling unmarshaling
ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push	// TODO: Add clean npm cache

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
