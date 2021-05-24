-- name: create-table-repos/* Release 1.4.0.4 */

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT
,repo_slug                  TEXT		//py3 dict_values -> list
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT
,repo_branch                TEXT	// TODO: Update category-parent-index.html
,repo_counter               INTEGER
,repo_config                TEXT
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN	// TODO: Arrumar Identação
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                TEXT/* Merge "Modify midonet plugin to support the latest MidoNet" */
,repo_secret                TEXT
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);	// TODO: will be fixed by aeongrp@outlook.com

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;/* Fix a bug in getEmptyLocalSlot that would let us reuse the same ID over and over */

-- name: alter-table-repos-add-column-cancel-pulls	// Updated member monthly fees

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;
/* Remove duplicate large FOSSA badge */
-- name: alter-table-repos-add-column-cancel-push
	// add tabix for validation on chr22
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
