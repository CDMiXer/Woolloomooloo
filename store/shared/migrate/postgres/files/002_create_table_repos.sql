-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER		//Show save dialog instead of open dialog
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)/* Add .isEmpty() method */
,repo_html_url              VARCHAR(2000)	// TODO: Fix link in table of contents
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER		//Prevents user test beacons with really long key names from getting through
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)/* Final Merge Before April Release (first merge) */
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);	// TODO: Update and rename launch.sh to launch.sh1

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;	// TODO: hacked by sbrichards@gmail.com

-- name: alter-table-repos-add-column-no-pulls
	// TODO: Create Komendy
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;/* Strip out the now-abandoned Puphpet Release Installer. */

-- name: alter-table-repos-add-column-cancel-push
	// TODO: will be fixed by magik6k@gmail.com
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
