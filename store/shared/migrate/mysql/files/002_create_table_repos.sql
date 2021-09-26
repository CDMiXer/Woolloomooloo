-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (/* TAG MetOfficeRelease-1.6.3 */
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)	// TODO: Adding in checks for comparing report IDs between samples
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)	// [ExoBundle] Initialization the form for the open question with one word
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)		//Merge "In integration tests wait 1 second after changing the password"
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
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
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER/* Fix response HTML formatting */
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork
	// fixing servlet passing parameters
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;
	// TODO: will be fixed by denner@gmail.com
-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;/* Delete object_script.vpropertyexplorer.Release */

-- name: alter-table-repos-add-column-cancel-pulls
	// TODO: Add dependant parameters
ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;
	// TODO: hacked by ligi@ligi.de
-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;/* Massively expanded content in README */
