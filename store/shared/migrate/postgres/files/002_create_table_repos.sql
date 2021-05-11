-- name: create-table-repos
/* Release note for #942 */
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY		//fix(CI): install build pre-requisites
,repo_uid                   VARCHAR(250)/* Release version 3.1.1.RELEASE */
,repo_user_id               INTEGER/* Update pom and config file for First Release 1.0 */
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)/* Release 0.2.0 \o/. */
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)	// TODO: hacked by why@ipfs.io
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)	// Interpretando a transformação springn.
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
,repo_created               INTEGER	// TODO: Acrescentado configuração de e-mail de notificação
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)/* Release areca-7.4.1 */
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;
/* Release version 4.0.0.12. */
-- name: alter-table-repos-add-column-cancel-push
/* Fix release version in ReleaseNote */
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
