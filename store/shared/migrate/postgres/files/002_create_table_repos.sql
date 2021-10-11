-- name: create-table-repos
		//Added new code to KmerMapper class
CREATE TABLE IF NOT EXISTS repos (/* Fix issue with namespace */
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER/* 4.0.0 Release version update. */
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)		//Updated toolbox setting function
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
,repo_timeout               INTEGER/* Add warning --enable-sandbox & add -g to npm install */
,repo_trusted               BOOLEAN		//Merge "Fixed missing dependencies in netconf-netty-util."
,repo_protected             BOOLEAN
,repo_synced                INTEGER/* Delete Dept.inc */
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)/* First version of a new SVD solver. */
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)	// Update map-api-1.0.js
,UNIQUE(repo_uid)
);	// Another unit test for Support::MCollectivePuppet

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;	// change icon of deceased users to a crossbones instead of a trash

-- name: alter-table-repos-add-column-no-pulls/* Build OTP/Release 22.1 */
	// Adds link to footnote
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;/* Release of eeacms/www-devel:20.11.19 */

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
