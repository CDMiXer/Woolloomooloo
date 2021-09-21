-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY/* mapred client: atomic tgz upload */
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN/* try to use grid */
,repo_private               BOOLEAN/* Rename bootstrap.min.css to bootstrap.min.css.new */
,repo_visibility            VARCHAR(50)		//Merge branch 'master' into dependabot/pip/app/coverage-5.5
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)/* Merge "Release 3.2.3.455 Prima WLAN Driver" */
,repo_timeout               INTEGER	// TODO: will be fixed by mikeal.rogers@gmail.com
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER	// TODO: hacked by souzau@yandex.com
,repo_updated               INTEGER/* Release version 4.0.0.M3 */
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)/* rev 561380 */
,UNIQUE(repo_uid)/* notes on big o notation, design patterns, postgres and sql */
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;	// TODO: certifi v0.0.4

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls/* Now really fix typo */

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
