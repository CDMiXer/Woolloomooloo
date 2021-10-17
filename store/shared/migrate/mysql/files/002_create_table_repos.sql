-- name: create-table-repos
/* Update title for live message */
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER	// Move constants and add count by level request
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)/* Release bzr-1.6rc3 */
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)	// TODO: hacked by jon@atack.com
,repo_active                BOOLEAN	// creacion de la ventana dialogo registro
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER/* Update How to use RemindMe.md */
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN		//Update dependencies for Symfony2.3 support
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)/* Properly disable in EditNode the property JSON Path when format BASE64 or HEX */
,UNIQUE(repo_uid)
);/* 2.0.7-beta5 Release */

-- name: alter-table-repos-add-column-no-fork	// Wrong form var

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls/* Better loading */

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push	// TODO: will be fixed by cory@protocol.ai

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;/* Release note 8.0.3 */
