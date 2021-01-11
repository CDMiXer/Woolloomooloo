-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)	// TODO: hacked by yuvalalaluf@gmail.com
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)/* Screenshots in Media Header Submenu fixxed ! */
,repo_slug                  VARCHAR(250)		//Delete .active_record_model_extension.rb.swp
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)	// Add a few templates
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)	// TODO: hacked by peterke@gmail.com
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN		//misc material: read wikipedia files and create w2v model
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);	// TODO: hacked by julia@jvns.ca

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;	// Number type enforced in loop

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;/* ignore stuff. */
