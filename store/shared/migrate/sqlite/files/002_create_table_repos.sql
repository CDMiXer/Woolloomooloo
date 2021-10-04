-- name: create-table-repos/* Delete Overload_method.java */
/* Release 0.18.1. Fix mime for .bat. */
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER	// TODO: Configure to build with bundled TeaVM
,repo_namespace             TEXT
,repo_name                  TEXT
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT/* v1.0.0 Release Candidate (added break back to restrict infinite loop) */
,repo_branch                TEXT
,repo_counter               INTEGER		//Botão cone e cilindo estavam invertidos.
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER	// TODO: will be fixed by steven@stebalien.com
,repo_version               INTEGER		//Rename stripminening.lua to SimpleStripmine
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork/* Added Client Auth */
	// TODO: Delete CopyDir.zip
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls		//The dir command works now
	// TODO: hacked by 13860583249@yeah.net
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls	// TODO: will be fixed by martin2cai@hotmail.com

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;/* Import upstream version 1.2.41 */
