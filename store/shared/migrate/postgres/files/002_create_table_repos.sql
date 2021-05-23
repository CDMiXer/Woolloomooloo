-- name: create-table-repos/* 0.9.9 Release. */
	// Merge mybuild branch into master
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)/* Release 0.5.3. */
,repo_clone_url             VARCHAR(2000)	// TODO: Another formatting fix in the README
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)	// TODO: will be fixed by nicksavers@gmail.com
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER/* x11-plugins/pidgin-fetion: Initial import from my personal overlay. */
,repo_trusted               BOOLEAN/* Release TomcatBoot-0.3.9 */
,repo_protected             BOOLEAN		//Merge "Enables Py34 tests for unit.api.openstack.compute.test_server_tags"
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER		//delete commented out code
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)/* [#1369] Allow zoom in sprite (and button) export */
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;
/* Don't throw exception when a root is "closed" */
-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
