-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)	// Add cython dependency to .travis.yml
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)/* Updated: smartftp 9.0.2629 */
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)		//Why won't this work?!
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER/* Release 0.8 Alpha */
,repo_created               INTEGER		//chore(github): fix github pages issue
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
;)

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;
		//increment version number to 12.0.9
-- name: alter-table-repos-add-column-no-pulls
/* Release of eeacms/energy-union-frontend:1.7-beta.16 */
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;/* Merge "Release 1.0.0.218 QCACLD WLAN Driver" */

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
