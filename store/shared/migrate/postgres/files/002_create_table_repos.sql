-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (/* Release 1.9 */
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)/* Release v2.3.0 */
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)		//Delete install_test_tools.sh
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)/* upload: early return when no files to upload */
,repo_scm                   VARCHAR(50)	// TODO: hacked by peterke@gmail.com
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN/* add self to servicename */
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)		//Corrected FC 16 buffer
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)/* Fix history */
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;
	// TODO: hacked by admin@multicoin.co
-- name: alter-table-repos-add-column-cancel-push/* Release 0.95.130 */

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
