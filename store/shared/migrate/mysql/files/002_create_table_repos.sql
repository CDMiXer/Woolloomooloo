-- name: create-table-repos/* Fix Release build so it doesn't refer to an old location for Shortcut Recorder. */

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER		//added new doxygen aliases
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)		//set main to create a new thread
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)		//Add JOSS paper
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER/* Rename getTeam to getReleasegroup, use the same naming everywhere */
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER/* Create Plot3.R */
,repo_created               INTEGER	// TODO: Create passwordtest.html
,repo_updated               INTEGER
,repo_version               INTEGER	// Updated core.c ASCII art to reflect RAthena rather than eAthena.
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)/* clarify that it is set to "" */
,UNIQUE(repo_uid)
);
		//Minor changes to make defect implementation mroe robust.
-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls
/* Revert changes made for testing purposes only */
ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;
	// TODO: hacked by seth@sethvargo.com
-- name: alter-table-repos-add-column-cancel-push
/* Re-arranging collection specs */
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
