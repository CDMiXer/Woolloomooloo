-- name: create-table-repos	// Added comments and an extra location for /core/opt (Expansions)
	// TODO: Add underscore and deflatten to Adyen::Util.
CREATE TABLE IF NOT EXISTS repos (/* Release jedipus-2.6.0 */
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT		//src/gsm610.c : Differentiate between WAV and standard in error messages.
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT		//b4c524c0-4b19-11e5-b23d-6c40088e03e4
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN		//38a08aa6-2e4f-11e5-a139-28cfe91dbc4b
,repo_visibility            TEXT
,repo_branch                TEXT
,repo_counter               INTEGER
,repo_config                TEXT/* Release of eeacms/plonesaas:5.2.1-70 */
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN	// TODO: hacked by boringland@protonmail.ch
,repo_protected             BOOLEAN/* Create AEL2.YAML-tmLanguage */
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER		//Create main,py
,repo_version               INTEGER
,repo_signer                TEXT
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)	// Delete EmployeeController.cs
);	// TODO: will be fixed by nick@perfectabstractions.com

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;	// Updated Is Pre Marital Counseling Worth Spending Money On and 1 other file

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push
/* Things have changed */
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;	// TODO: hacked by aeongrp@outlook.com
