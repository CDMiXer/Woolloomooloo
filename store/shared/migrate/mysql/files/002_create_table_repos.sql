-- name: create-table-repos
	// TODO: Delete datapath_tb.v
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)	// Merge "Cancel a screenshot notification after a share target is chosen."
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)		//Create pdf.svg
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)/* Release Notes for v02-15-02 */
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER		//pretty printed
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER/* Release 1.061 */
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)/* Merge "Do not register more than one panic for a single recipe." into develop */
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)	// TODO: 897b92c2-2e63-11e5-9284-b827eb9e62be
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork
		//removed ES
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;
		//Added IMDB module
-- name: alter-table-repos-add-column-no-pulls		//this is convore not tablib; thought i would get some low hanging fruit. 
/* Release Notes: Q tag is not supported by linuxdoc (#389) */
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;
		//Initial implementation of property editor
-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push
	// TODO: will be fixed by steven@stebalien.com
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;	// Tagging checker-272.
