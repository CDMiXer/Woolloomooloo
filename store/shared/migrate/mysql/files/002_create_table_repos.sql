-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)		//3ec9bfe2-2e5e-11e5-9284-b827eb9e62be
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)		//flowtype.js added
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)/* Update references to removed method in main() */
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)/* [4723] changed jetty runlevel to 4, after jpa is started on 3 */
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER	// TODO: will be fixed by why@ipfs.io
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)/* fixed small formating error */
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);/* Merge "Improve ImageView drawable re-use" into mnc-dev */

-- name: alter-table-repos-add-column-no-fork/* Some pyx optimizations */

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;
		//Update Lab 6.md
-- name: alter-table-repos-add-column-no-pulls
		//Remove sorting on searched#index for packages
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
