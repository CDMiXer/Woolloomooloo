-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (	// TODO: will be fixed by why@ipfs.io
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)	// TODO: Added @dave91087
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN/* Edit. readme */
,repo_synced                INTEGER
,repo_created               INTEGER/* End html tag off */
,repo_updated               INTEGER
,repo_version               INTEGER/* Updating the projects timeline */
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;		//Delete automation05.py~

-- name: alter-table-repos-add-column-no-pulls/* split embed_source method up to be more clear */

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls/* Release v0.5.8 */

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push
	// TODO: hacked by mowrain@yandex.com
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
