-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER/* Release of eeacms/forests-frontend:1.8-beta.3 */
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)		//Create MAEC 5.0 index.md page
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN/* Merge "[doc] Release Victoria" */
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER	// TODO: will be fixed by aeongrp@outlook.com
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);		//system/fd_util: remove open_cloexec()

-- name: alter-table-repos-add-column-no-fork
/* Release Ver. 1.5.3 */
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;	// TODO: Implement ImageFactory for images

-- name: alter-table-repos-add-column-no-pulls
/* e51cb06e-327f-11e5-bb0e-9cf387a8033e */
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
