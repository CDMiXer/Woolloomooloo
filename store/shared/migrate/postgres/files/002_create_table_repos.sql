-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (	// TODO: changed version-string to <0.4.1-testing>
 repo_id                    SERIAL PRIMARY KEY/* Removed unused code; fixed some docs; added dependency */
,repo_uid                   VARCHAR(250)		//Add warning for missing s in rwatershedflags.
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)		//WIP: Added Support for loading config files in .coffee or .js format 
,repo_clone_url             VARCHAR(2000)/* Publishing post - a simple bash script */
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN		//Changed heading and allow ordering by system
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER/* Outline manual */
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER	// TODO: Update clearmap-spotdetection.md
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN/* Added Release Notes for v0.9.0 */
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)/* captcha antiguo contacto quitado */
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)	// Add BFKit & BFKit-Swift to Utility section
);

-- name: alter-table-repos-add-column-no-fork	// TODO: Add NewQuery and generic ConfigureQuery

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;
		//37d5d65a-2e48-11e5-9284-b827eb9e62be
-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
