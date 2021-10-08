-- name: create-table-repos/* Added Apache Commons Java libraries to the repository's java lib. */

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT	// TODO: New version of Cherish - 1.1
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER/* Release dhcpcd-6.10.2 */
,repo_namespace             VARCHAR(250)/* Release notes for 0.1.2. */
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER		//initial change
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN/* back to jimhester wercker */
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)	// TODO: hacked by steven@stebalien.com
,UNIQUE(repo_slug)	// TODO: will be fixed by cory@protocol.ai
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork
	// TODO: boot.jar doesn't have to be included in Frameworks
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls		//Lock update process and put cache in data folder

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;/* Merge "Fix a quoting typo" */

-- name: alter-table-repos-add-column-cancel-pulls
/* Use new rake task on Travis */
ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;	// Updated the gperf feedstock.

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
